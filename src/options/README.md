# [Fancy Settings 1.2](https://github.com/frankkohlhepp/fancy-settings)
*Create fancy, chrome-look-alike settings for your Chrome or Safari extension in minutes!*

### Howto
Welcome to Fancy Settings! Are you ready for tabs, groups, search, good style?  
Let's get started, it only takes a few minutes...

Settings can be of different types: text input, checkbox, slider, etc. Some "settings" are not actual settings but provide functionality that is relevant to the options page: description (which is simply a block of text), button.

Settings are defined in the manifest.js file as JavaScript objects. Each setting is defined by specifying a number of parameters. All types of settings are configured with the string parameters tab, group, name and type.

###Basic example:
```javascript
{
    "tab": "Tab 1",
    "group": "Group 1",
    "name": "checkbox1",
    "type": "checkbox"
}
```

"name" is used as a part of the key when storing the setting's value in localStorage. 
If it's missing, nothing will be saved.

###Additionally, all types of settings are configured with their own custom parameters:

###Description ("type": "description")

text (string) the block of text, which can include HTML tags. You can continue multiple lines of text by putting a \ at the end of a line, just as with any JavaScript file.

####
Button ("type": "button")
```
 Label (string) text shown in front of the button

 Text (string) text shown on the button
```

####Text ("type": "text")
```
 label (string) text shown in front of the text field

 text (string) text shown in the text field when empty

 masked (boolean) indicates a password field
```

####Checkbox ("type": "checkbox")
```
 label (string) text shown behind the checkbox
```

####Slider ("type": "slider")
```
 label (string) text shown in front of the slider

 max (number) maximal value of the slider

 min (number) minimal value of the slider 

 step (number) steps between two values

 display (boolean) indicates whether to show the slider display

 displayModifier (function) a function to modify the value shown in the display
```

####PopupButton ("type": "popupButton"), ListBox ("type": "listBox") & RadioButtons ("type": "radioButtons")
```
label (string) text shown in front of the options

 options (array of options)

 where an option can be one of the following formats:
```

####"value"
```
["value", "displayed text"]

{value: "value", text: "displayed text"}
```
The "displayed text" field is optional and is displayed to the user when you don't want to display the internal value directly to the user.

#### You can also group options so that the user can easily choose among them (groups may only be applied to popupButtons):

```javascript
          "options": {
              "groups": [
                  "Hot", "Cold",
              ],
              "values": [
                  {
                      "value": "hot",
                      "text": "Very hot",
                      "group": "Hot",
                  },
                  {
                      "value": "Medium",
                      "group": 1,
                  },
                  {
                      "value": "Cold",
                      "group": 2,
                  },
                  ["Non-existing"]
              ],
          },

```

### License
Fancy Settings is licensed under the **LGPL 2.1**.  
For details see *LICENSE.txt*