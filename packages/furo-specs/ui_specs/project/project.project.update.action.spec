{
  "_writeprotection": false,
  "theme": "ActionBaseTheme",
  "class_name": "ProjectProjectUpdateAction",
  "component_name": "project-project-update-action",
  "description": "service specs for the project api",
  "source": "./ui_specs/project/project.project.update.action.spec",
  "service_name": "ProjectService",
  "response_type": "project.ProjectEntity",
  "imports": [],
  "items": [
    {
      "label": "save",
      "rel": "update",
      "icon": null,
      "component": "furo-button",
      "onclick": "-^update-req",
      "flags": [
        "primary",
        "unelevated"
      ],
      "attrs": []
    },
    {
      "label": "reload",
      "rel": "self",
      "icon": null,
      "component": "furo-button",
      "onclick": "-^self-req",
      "flags": [
        "outline"
      ],
      "attrs": []
    },
    {
      "component": "furo-empty-spacer"
    },
    {
      "label": "cancel",
      "rel": "",
      "icon": null,
      "component": "furo-button",
      "onclick": "-^reset-req",
      "flags": [
        "outline"
      ],
      "attrs": []
    },
    {
      "label": "delete",
      "rel": "delete",
      "icon": "delete",
      "component": "furo-button",
      "onclick": "-^delete-req",
      "flags": [
        "unelevated",
        "danger"
      ],
      "attrs": []
    }
  ],
  "request_type": "project.Project"
}