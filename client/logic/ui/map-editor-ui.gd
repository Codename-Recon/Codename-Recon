@tool
extends Control

signal select_terrain(terrain_set: int, terrain: int)
signal resize_map(new_size: Vector2i)

@export var tile_set: TileSet
@onready var main_menu: Panel = $MainMenu
@onready var terrains_container: HBoxContainer = $Panel/TerrainsContainer
@onready var resize_menu: Panel = $"ResizeMenu"
var last_button: Button


func _ready() -> void:
	if tile_set:
		for child: Node in terrains_container.get_children():
			terrains_container.remove_child(child)
		for terrain_idx: int in range(tile_set.get_terrains_count(0)):
			var button: Button = Button.new()
			button.text = tile_set.get_terrain_name(0, terrain_idx)
			button.pressed.connect(func() -> void: on_terrain_selected(button, terrain_idx))
			terrains_container.add_child(button)


func on_terrain_selected(button: Button, terrain_idx: int) -> void :
	select_terrain.emit(0, terrain_idx)
	if last_button:
		last_button.disabled = false
	last_button = button
	last_button.disabled = true

func _on_close_menu_pressed() -> void:
	main_menu.hide()

func _on_menu_pressed() -> void:
	main_menu.show()

func _on_resize_pressed() -> void:
	resize_menu.show()
	main_menu.hide()
