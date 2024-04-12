@icon("res://assets/images/icons/nodes/layers-outline.svg")
@tool
class_name Map
extends Node2D

const DUPLICATE_TEST_SIZE: int = 4
const TERRAIN_STATS: PackedScene = preload("res://logic/ui/terrain_stats.tscn")

@export var map_name: String
@export var creator: String
@export var creator_url: String
@export_group("Tool")
@export var test_duplicates: bool:
	set(value):
		_test_for_duplicates()
@export_multiline var duplicate_result: String = ""

var players: Node

static var predefined_terrains_packed_scenes: Dictionary:
	get:
		if not predefined_terrains_packed_scenes:
			predefined_terrains_packed_scenes = _load_predefined_terrains_packed_scenes(_terrain_path)
		return predefined_terrains_packed_scenes
static var predefined_units_packed_scenes: Dictionary:
	get:
		if not predefined_units_packed_scenes:
			predefined_units_packed_scenes = _load_predefined_units_packed_scenes(_unit_path)
			print()
		return predefined_units_packed_scenes

static var _terrain_path: String = "res://logic/game/terrain/"
static var _unit_path: String = "res://logic/game/units/"

var _types: GlobalTypes = Types
var _players_node: Node

static func _load_predefined_terrains_packed_scenes(path: String) -> Dictionary:
	var _predefines: Dictionary = {}
	var dir: DirAccess = DirAccess.open(path)
	if not dir:
		push_error("Can not open directory: " + path)
	for file_name: String in dir.get_files():
		if file_name.ends_with(".tscn.remap"):
			file_name = file_name.trim_suffix(".remap")
		if not file_name.ends_with(".tscn"):
			continue
		var terrain_packed_scene: PackedScene = load(path + file_name) as PackedScene
		var terrain: Terrain = terrain_packed_scene.instantiate()
		_predefines[terrain.tile_id] = terrain_packed_scene
		terrain.queue_free()
	return _predefines
	
	
static func _load_predefined_units_packed_scenes(path: String) -> Dictionary:
	var _predefines: Dictionary = {}
	var dir: DirAccess = DirAccess.open(path)
	if not dir:
		push_error("Can not open directory: " + path)
	for file_name: String in dir.get_files():
		if file_name.ends_with(".tscn.remap"):
			file_name = file_name.trim_suffix(".remap")
		if not file_name.ends_with(".tscn"):
			continue
		var unit_packed_scene: PackedScene = (load(path + file_name) as PackedScene)
		var unit: Unit = unit_packed_scene.instantiate()
		_predefines[unit.id] = unit_packed_scene
		unit.queue_free()
	return _predefines


func has_player(player_id: int) -> bool:
	return players.get_children().any(func(p: Player) -> bool: return p.id == player_id)


func has_terrain_or_unit_owned_by_player(player_id: int) -> bool:
	if not has_player(player_id):
		return false
	for c: Node in get_children():
		if c is Terrain:
			var terrain: Terrain = c
			if terrain.player_owned and terrain.player_owned.id == player_id:
				return true
			if terrain.has_unit():
				if terrain.get_unit().player_owned.id == player_id:
					return true
	return false


func create_or_get_player(player_id: int) -> Player:
	if player_id == 0:
		return null
	if has_player(player_id):
		return get_player(player_id)
	var player: Player = Player.new()
	player.id = player_id
	player.name = "PLAYER %s" % player_id
	players.add_child(player)
	return player


func get_player(player_id: int) -> Player:
	if not has_player(player_id) or player_id == 0:
		return null
	return players.get_children().filter(func(p: Player) -> bool: return p.id == player_id)[0]


func remove_player(player_id: int) -> void:
	if not has_player(player_id):
		return
	var player: Player = get_player(player_id)
	player.queue_free()


## Creats a terrain and updates the players
func create_terrain(id: String, tile_id: String, terrain_position: Vector2i, texture: Texture2D, ground_tile_texture: Texture2D, player_id: int) -> void:
	# Check if predefined terrain exist. If not -> create terrain
	var terrain: Terrain
	if tile_id in Map.predefined_terrains_packed_scenes:
		var terrain_scene: PackedScene = Map.predefined_terrains_packed_scenes[tile_id]
		terrain = terrain_scene.instantiate()
	else:
		if not texture:
			push_error("Can not create terrain %s without texture " % id)
			return
		terrain = Terrain.new()
		terrain.id = id
		terrain.tile_id = tile_id
		var sprite: Sprite2D = Sprite2D.new()
		sprite.name = "Sprite2D"
		terrain.add_child(sprite)
		terrain.sprite = sprite
		terrain.sprite.texture = texture
	# Add ground tile if specified in the tile set
	if ground_tile_texture:
		var ground_sprite: Sprite2D = Sprite2D.new()
		ground_sprite.name = "Ground"
		terrain.add_child(ground_sprite)
		terrain.move_child(ground_sprite, 0)
		ground_sprite.texture = ground_tile_texture
	# Add terrain stats
	terrain.add_child(TERRAIN_STATS.instantiate())
	# Add player
	if _types.terrains[id]["can_capture"]:
		var player: Player = create_or_get_player(player_id)
		terrain.player_owned = player
	add_child(terrain)
	terrain.name = id
	terrain.position = terrain_position
	
	
## Creats a unit on a specific terrain and updates the players. If a unit is already on that terrain, it will be replaced with the new unit.
## The return value indicates whether the unit can be placed at that specific location.
func create_unit(id: String, terrain_position: Vector2i, player_id: int) -> bool:
	var tmp_terrain: Terrain = get_tree().get_nodes_in_group("terrain")[0]
	var terrain: Terrain = tmp_terrain.get_terrain_by_position(terrain_position)
	if terrain:
		if terrain.has_unit():
			remove_unit(terrain.position)
		var unit_packed_scene: PackedScene = Map.predefined_units_packed_scenes[id]
		var unit: Unit = unit_packed_scene.instantiate()
		var movement_type: String = _types.units[unit.id]["movement_type"]
		var movement_value: int = _types.movements[terrain.id]["CLEAR"][movement_type]
		# If unit is not allowed to be placed here
		if movement_value == 0:
			unit.queue_free()
			return false
		# Add player
		var player: Player = create_or_get_player(player_id)
		unit.player_owned = player
		terrain.add_child(unit)
		return true
	return false


## Removes a terrain and updates the players
func remove_terrain(terrain_position: Vector2i) -> void:
	var tmp_terrain: Terrain = get_tree().get_nodes_in_group("terrain")[0]
	var terrain: Terrain = tmp_terrain.get_terrain_by_position(terrain_position)
	if terrain:
		# Remove terrain and also remove player if no other entities are owned by it
		var player: Player = terrain.player_owned
		remove_child(terrain)
		terrain.queue_free()
		if player and not has_terrain_or_unit_owned_by_player(player.id):
			remove_player(player.id)


## Removes a unit and updates the players
func remove_unit(unit_position: Vector2i) -> void:
	var tmp_terrain: Terrain = get_tree().get_nodes_in_group("terrain")[0]
	var terrain: Terrain = tmp_terrain.get_terrain_by_position(unit_position)
	if terrain and terrain.has_unit():
		# Remove terrain and also remove player if no other entities are owned by it
		var unit: Unit = terrain.get_unit()
		var player: Player = unit.player_owned
		terrain.remove_child(unit)
		unit.queue_free()
		if player and not has_terrain_or_unit_owned_by_player(player.id):
			remove_player(player.id)
	
	
func _ready() -> void:
	if not Engine.is_editor_hint():
		if not has_node("Players"):
			_players_node = Node.new()
			_players_node.name = "Players"
			add_child(_players_node)
			players = _players_node
		else:
			players = get_node("Players")


func _test_for_duplicates() -> void:
	var found_duplicates: bool = false
	duplicate_result = ""
	for child: Node in get_children():
		if is_instance_of(child, Terrain):
			var terrain: Terrain = child as Terrain
			var result: Array[Node] = get_children().filter(func(x: Node) -> bool: 
				var length_squared: float = ((x as Terrain).global_position - terrain.global_position).length_squared()
				return x != terrain and is_instance_of(x, Terrain)	and length_squared < DUPLICATE_TEST_SIZE)
			for r: Terrain in result:
				duplicate_result += str(r) + "\n"
				found_duplicates = true
	if not found_duplicates:
		duplicate_result = "No duplicates found (%s)" % Time.get_datetime_string_from_system()
