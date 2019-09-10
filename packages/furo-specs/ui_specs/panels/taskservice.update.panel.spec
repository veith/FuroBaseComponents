{
  "class_name": "TaskServiceUpdatePanel",
  "component_name": "taskservice-update-panel",
  "description": "Updates a Task, partial updates are not supported",
  "source": "./specs/task/task.service.spec",
  "imports": [
    "../forms/task-task-form",
    "../actions/task-task-update-action"
  ],
  "form": {
    "name": "task-task-form",
    "attrs": []
  },
  "action": {
    "name": "task-task-update-action",
    "listen_to": [
      "update"
    ],
    "attrs": []
  }
}