
# image-magick 1.0

Image modifications in Direktiv

---
- #### Categories: tools
- #### Image: gcr.io/direktiv/functions/image-magick 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/image-magick/issues
- #### URL: https://github.com/direktiv-apps/image-magick
- #### Maintainer: [direktiv.io](https://www.direktiv.io) 
---

## About image-magick

This app can run multiple [Image Magick](https://imagemagick.org/index.php) commands.  The results can either be stored in the output folder of Direktiv to store them as variables or returned as base64.

### Example(s)
  #### Function Configuration
```yaml
functions:
- id: image-magick
  image: gcr.io/direktiv/functions/image-magick:1.0
  type: knative-workflow
```
   #### Basic
```yaml
- id: image-magick
  type: action
  action:
    function: image-magick
    input: 
      commands:
      - command: 'convert mypic.png json:'
      - command: convert mypic.png -fuzz 25% -fill red -opaque white -flatten mypic.png
      - command: convert mypic.png -resize 200x100 mypic.jpg
  catch:
  - error: '*'
```
   #### Advanced
```yaml
- id: modify 
  type: action
  action:
    function: image-magick
    input: 
      commands:
      # stores the image in workflow scope variable `resized.png`
      - convert mypic.png -resize 200% out/workflow/resized.png 
  catch: 
  - error: "*"
```

   ### Secrets


*No secrets required*







### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  List of executed commands.
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
[
  {
    "result": null,
    "success": true
  },
  {
    "result": null,
    "success": true
  }
]
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| image-magick | [][PostOKBodyImageMagickItems](#post-o-k-body-image-magick-items)| `[]*PostOKBodyImageMagickItems` |  | |  |  |
| images | [][PostOKBodyImagesItems](#post-o-k-body-images-items)| `[]*PostOKBodyImagesItems` |  | |  |  |


#### <span id="post-o-k-body-image-magick-items"></span> postOKBodyImageMagickItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-o-k-body-images-items"></span> postOKBodyImagesItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | `[{"command":"echo Hello"}]`| Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| return | []string| `[]string` |  | | Returns the images as base64 | `myimage.jpg` |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run |  |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
