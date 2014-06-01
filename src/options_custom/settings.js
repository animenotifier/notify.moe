window.addEvent("domready", function () {
	// Option 1: Use the manifest:
	new FancySettings.initWithManifest(function (settings) {
		var updateURL = function() {
			var url = animeListProviders[settings.manifest.animeListProvider.element.value].url + settings.manifest.userName.element.value;
			settings.manifest.previewURL.element.innerHTML = "<a href='" + url + "' target='_blank'>" + url + "</a>";
		};

		settings.manifest.userName.addEvent("action", updateURL);
		settings.manifest.animeListProvider.addEvent("action", updateURL);

		if(!("store.settings.updateInterval" in localStorage))
			settings.manifest.updateInterval.set("5");

        if(!("store.settings.maxEpisodeDifference" in localStorage))
            settings.manifest.maxEpisodeDifference.set("1");

		updateURL();
        
        for(var key in localStorage) {
            if(key.substr(0, 6) != "anime.")
                continue;

            var title = key.substr(6);
            //var anime = localStorage.getObject(key);

            var createSettings = function() {
	            animeUpdater.queryPossibleAnimeOptions(title, "", function(animeTitle, qualityOptions, subsOptions) {
	            	var customSearchTitle = settings.create({
	                    "tab": "Filters",
	                    "group": animeTitle,
	                    "name": animeTitle + ":search",
	                    "type": "text",
	                    "label": "Custom search title:"
	                });

	                var subs = settings.create({
	                    "tab": "Filters",
	                    "group": animeTitle,
	                    "name": animeTitle + ":subs",
	                    "type": "popupButton",
	                    "label": "Subs:",
	                    "options": subsOptions
	                });

	                var qualities = settings.create({
	                    "tab": "Filters",
	                    "group": animeTitle,
	                    "name": animeTitle + ":quality",
	                    "type": "listBox",
	                    "label": "Quality:",
	                    "options": qualityOptions
	                });

	                var updateQualityOptions = function() {
	                    var qualitiesParent = qualities;

	                    animeUpdater.queryPossibleAnimeOptions(animeTitle, subs.get(), function(animeTitle, newQualityOptions, subsOptions) {
	                        qualities.element.innerHTML = "";
	                        var key = "store.settings." + animeTitle + ":quality";
	                        var selectedQuality = localStorage[key];
	                        selectedQuality = selectedQuality ? selectedQuality.replace(/"/g, "") : "";
	                        var optionFound = false;

	                        for(var i = 0; i < newQualityOptions.length; i++) {
	                            var optionInfo = newQualityOptions[i];
	                            var option = document.createElement("option");
	                            var value = optionInfo["value"];

	                            if(value == selectedQuality)
	                                optionFound = true;

	                            option.setAttribute("value", value);
	                            option.text = optionInfo["text"];
	                            
	                            qualitiesParent.element.add(option);
	                        }

	                        if(optionFound) {
	                            qualitiesParent.element.value = selectedQuality;
	                            localStorage[key] = selectedQuality;
	                        } else {
	                            qualitiesParent.element.value = "";
	                            localStorage[key] = "";
	                        }
	                    });
	                };

	                //customSearchTitle.addEvent("action", updateSubs);
	                subs.addEvent("action", updateQualityOptions);
	                updateQualityOptions();

	                // Align
	                // TODO: Fix this
	                var animeSettings = [];

	                animeSettings.push(subs);
	                animeSettings.push(qualities);
	                animeSettings.push(customSearchTitle);

	                settings.align(animeSettings);
	            });
			};

			createSettings();
        }
	});
	
	// Option 2: Do everything manually:
	/*
	var settings = new FancySettings("My Extension", "icon.png");
	
	var username = settings.create({
		"tab": i18n.get("information"),
		"group": i18n.get("login"),
		"name": "username",
		"type": "text",
		"label": i18n.get("username"),
		"text": i18n.get("x-characters")
	});
	
	var password = settings.create({
		"tab": i18n.get("information"),
		"group": i18n.get("login"),
		"name": "password",
		"type": "text",
		"label": i18n.get("password"),
		"text": i18n.get("x-characters-pw"),
		"masked": true
	});
	
	var myDescription = settings.create({
		"tab": i18n.get("information"),
		"group": i18n.get("login"),
		"name": "myDescription",
		"type": "description",
		"text": i18n.get("description")
	});
	
	var myButton = settings.create({
		"tab": "Information",
		"group": "Logout",
		"name": "myButton",
		"type": "button",
		"label": "Disconnect:",
		"text": "Logout"
	});
	
	// ...
	
	myButton.addEvent("action", function () {
		alert("You clicked me!");
	});
	
	settings.align([
		username,
		password
	]);
	*/
});
