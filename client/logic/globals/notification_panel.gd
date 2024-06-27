extends Panel

class_name NotificationPanel 

@export
var title: String : 
	set(new_text):
		title = new_text
		if title_label:
			update_title()

@export
var actions: Array[Button]:
	set(new_actions):
		actions = new_actions
		if actions_container:
			update_actions()
			
@export
var timeout: int


@onready
var title_label: Label = $Title

@onready
var actions_container: HBoxContainer = $Actions

@onready
var timer: Timer = $Timer

func _ready() -> void:
	update_title()
	update_actions()
	if timeout>0:
		timer.wait_time = float(timeout)
		timer.start()

func update_title() -> void:
	title_label.text = title

func update_actions() -> void:
	for child: Node in actions_container.get_children():
		actions_container.remove_child(child)
	for action: Button in actions:
		actions_container.add_child(action)


func _on_close_pressed() -> void:
	queue_free();
