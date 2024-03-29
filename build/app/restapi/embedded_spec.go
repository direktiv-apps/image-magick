// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Image modifications in Direktiv",
    "title": "image-magick",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "tools"
      ],
      "container": "gcr.io/direktiv/functions/image-magick",
      "issues": "https://github.com/direktiv-apps/image-magick/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This app can run multiple [Image Magick](https://imagemagick.org/index.php) commands.  The results can either be stored in the output folder of Direktiv to store them as variables or returned as base64.",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/image-magick"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "default": [
                    {
                      "command": "echo Hello"
                    }
                  ],
                  "items": {
                    "type": "object",
                    "properties": {
                      "command": {
                        "description": "Command to run",
                        "type": "string"
                      },
                      "continue": {
                        "description": "Stops excecution if command fails, otherwise proceeds with next command",
                        "type": "boolean"
                      },
                      "print": {
                        "description": "If set to false the command will not print the full command with arguments to logs.",
                        "type": "boolean",
                        "default": true
                      },
                      "silent": {
                        "description": "If set to false the command will not print output to logs.",
                        "type": "boolean",
                        "default": false
                      }
                    }
                  }
                },
                "files": {
                  "description": "File to create before running commands.",
                  "type": "array",
                  "default": null,
                  "items": {
                    "$ref": "#/definitions/direktivFile"
                  }
                },
                "return": {
                  "description": "Returns the images as base64",
                  "type": "array",
                  "items": {
                    "type": "string"
                  },
                  "example": "myimage.jpg"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "properties": {
                "image-magick": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "success",
                      "result"
                    ],
                    "properties": {
                      "result": {
                        "additionalProperties": false
                      },
                      "success": {
                        "type": "boolean"
                      }
                    }
                  }
                },
                "images": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "success",
                      "result"
                    ],
                    "properties": {
                      "result": {
                        "additionalProperties": false
                      },
                      "success": {
                        "type": "boolean"
                      }
                    }
                  }
                }
              }
            },
            "examples": {
              "image-magick": [
                {
                  "result": null,
                  "success": true
                },
                {
                  "result": null,
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            },
            {
              "action": "foreach",
              "exec": "base64 -w 0 {{ .Item }}",
              "loop": ".Return"
            }
          ],
          "output": "{\n  \"image-magick\": {{ index . 0 | toJson }}\n  {{ $l := len (index . 1) }}\n  {{- if gt $l 0 }}\n  , \"images\": {{ index . 1 | toJson }}\n  {{- end }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: image-magick\n  type: action\n  action:\n    function: image-magick\n    input: \n      commands:\n      - command: 'convert mypic.png json:'\n      - command: convert mypic.png -fuzz 25% -fill red -opaque white -flatten mypic.png\n      - command: convert mypic.png -resize 200x100 mypic.jpg\n  catch:\n  - error: '*'",
            "title": "Basic"
          },
          {
            "content": "- id: modify \n  type: action\n  action:\n    function: image-magick\n    input: \n      commands:\n      # stores the image in workflow scope variable ` + "`" + `resized.png` + "`" + `\n      - convert mypic.png -resize 200% out/workflow/resized.png \n  catch: \n  - error: \"*\"",
            "title": "Advanced"
          }
        ],
        "x-direktiv-function": "functions:\n- id: image-magick\n  image: gcr.io/direktiv/functions/image-magick:1.0\n  type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Image modifications in Direktiv",
    "title": "image-magick",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "tools"
      ],
      "container": "gcr.io/direktiv/functions/image-magick",
      "issues": "https://github.com/direktiv-apps/image-magick/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This app can run multiple [Image Magick](https://imagemagick.org/index.php) commands.  The results can either be stored in the output folder of Direktiv to store them as variables or returned as base64.",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/image-magick"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "image-magick": [
                {
                  "result": null,
                  "success": true
                },
                {
                  "result": null,
                  "success": true
                }
              ]
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            },
            {
              "action": "foreach",
              "exec": "base64 -w 0 {{ .Item }}",
              "loop": ".Return"
            }
          ],
          "output": "{\n  \"image-magick\": {{ index . 0 | toJson }}\n  {{ $l := len (index . 1) }}\n  {{- if gt $l 0 }}\n  , \"images\": {{ index . 1 | toJson }}\n  {{- end }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: image-magick\n  type: action\n  action:\n    function: image-magick\n    input: \n      commands:\n      - command: 'convert mypic.png json:'\n      - command: convert mypic.png -fuzz 25% -fill red -opaque white -flatten mypic.png\n      - command: convert mypic.png -resize 200x100 mypic.jpg\n  catch:\n  - error: '*'",
            "title": "Basic"
          },
          {
            "content": "- id: modify \n  type: action\n  action:\n    function: image-magick\n    input: \n      commands:\n      # stores the image in workflow scope variable ` + "`" + `resized.png` + "`" + `\n      - convert mypic.png -resize 200% out/workflow/resized.png \n  catch: \n  - error: \"*\"",
            "title": "Advanced"
          }
        ],
        "x-direktiv-function": "functions:\n- id: image-magick\n  image: gcr.io/direktiv/functions/image-magick:1.0\n  type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "image-magick": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/postOKBodyImageMagickItems"
          }
        },
        "images": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/postOKBodyImagesItems"
          }
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodyImageMagickItems": {
      "type": "object",
      "required": [
        "success",
        "result"
      ],
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodyImagesItems": {
      "type": "object",
      "required": [
        "success",
        "result"
      ],
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "properties": {
        "commands": {
          "description": "Array of commands.",
          "type": "array",
          "default": [
            {
              "command": "echo Hello"
            }
          ],
          "items": {
            "$ref": "#/definitions/postParamsBodyCommandsItems"
          }
        },
        "files": {
          "description": "File to create before running commands.",
          "type": "array",
          "default": [],
          "items": {
            "$ref": "#/definitions/direktivFile"
          }
        },
        "return": {
          "description": "Returns the images as base64",
          "type": "array",
          "items": {
            "type": "string"
          },
          "example": "myimage.jpg"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyCommandsItems": {
      "type": "object",
      "properties": {
        "command": {
          "description": "Command to run",
          "type": "string"
        },
        "continue": {
          "description": "Stops excecution if command fails, otherwise proceeds with next command",
          "type": "boolean"
        },
        "print": {
          "description": "If set to false the command will not print the full command with arguments to logs.",
          "type": "boolean",
          "default": true
        },
        "silent": {
          "description": "If set to false the command will not print output to logs.",
          "type": "boolean",
          "default": false
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
