# Anime Notifier

## What kind of website is this?

An anime tracker where you can add anime to your list and edit your episode progress using either the website, the chrome extension or the mobile app.

## Why is it called notify.moe?

Because we made a notifier that takes your watching list, checks it against external websites (currently twist.moe) and notifies you when there is a new episode on that external site.

## So it's just a notifier?

In the past it was, but we're growing bigger by establishing a database that combines information from multiple sources and also growing as a community. Many of us are hanging out on Discord and you are welcome to join us. We also have our own anime lists now due to popular request of adding episode progress changes to our browser extension.

## How do I use the search?

Press the "F" key while you're browsing the site and start searching for an anime title.

## How do I add an anime to my list?

Once you open the anime page you should see a button called "Add to my collection". Clicking that will add the anime to your "Plan to watch" list. To move it to your current "Watching" list, you need to click "Edit in collection" and change the status to "Watching".

## How do I edit my episode progress?

There are 2 ways of editing your progress:

1. Click on the "+" button that shows up when you hover over the episode number. This will increase your progress by one episode on each click.
1. Click on the episode number so that a text input cursor shows up. Use backspace/delete keys and enter your new number directly. Press Enter or click somewhere else to confirm.

## How do I edit my rating?

Your "Overall" rating can be edited with the same method as episodes by clicking on the number directly so that the text input cursor shows up, then entering a new number and confirming with Enter. The other 3 rating numbers for Story, Visuals and Soundtrack can only be edited by going into edit mode (click on the anime title in your list).

## How does the rating system work?

You can rate each entry in your anime list in 4 different categories:

* Overall (this will determine the sorting order)
* Story (how interesting was the story/plot?)
* Visuals (art & effect & animation quality)
* Soundtrack (music rating)

Each rating is a number on a scale of 0 to 10. A rating of 0 counts as "not rated" and will be ignored in average rating calculations for that anime. Thus the lowest possible rating you can assign to an anime is 0.1. The highest possible rating is 10. The average is close to the number 5.

## What are the community rules for conversations on the forum?

* Be respectful to each other.
* Realize that every person has his or her own opinion and that you should treat that opinion with respect. You do not have to agree with strangers on the internet but it's worth thinking about their viewpoint.
* Do not spam.
* Do not advertise unrelated products. If anything it needs to be related to anime or the site itself.
* We absolutely do not mind links to competitors or similar websites. Feel free to post them.

## How do I import my former anime list?

We added importers for what we consider to be the 3 most popular list providers:

* anilist.co
* kitsu.io
* myanimelist.net

To use an importer, enter your nickname for the site you want to import from and click the "Import" button with the list provider name that just appeared.

## How do I install the site as an Android App?

This website uses a very recent and modern technology that allows you to install websites as local apps. To install notify.moe as a mobile app, do the following:

1. Go to https://notify.moe on your Android device.
2. Open the menu by tapping the top right part of your browser.
3. Choose "Add to Home screen" and confirm it.
4. Now you can access your anime list easily from your home screen and get notified about new episodes.

## How do I install the site as a PC/Desktop App?

In Chrome, open the top right menu and go to **More tools > Add to desktop**.

## How do notifications work from a technical perspective?

There are many, many ways how notifications can be implemented from a technical standpoint. There is e.g. "polling" which means that an app periodically checks external sites and tells you when something new is available. We are not using polling because these periodic checks can quickly drain your battery on a mobile phone. We are using so-called "push notifications" instead. The advantage of push notifications is that your mobile phone or desktop PC doesn't have to do periodic checks anymore - instead the website will send new episode releases to all of your registered devices. This consumes less CPU/network resources and is much more battery friendly for mobile devices.

## Can you tell me more about the history of this software?

From a technological standpoint we went through quite a few different approaches:

* Version 1.0: This version was just a browser extension with **client-side JS**.
* Version 2.0: To decrease the number of requests/pressure on external sites we made a central website. It was written in **PHP**.
* Version 3.0: A complete remake of the website in **node.js** supporting 4 different list providers and 2 anime providers. Episode changes were not possible.
* Version 4.0: We switched to our own hosted anime lists to make episode updates in the extension as smooth as possible. The website is now written in **Go** and uses 3 separate servers/machines (web server, database and the scheduler).

## How many developers are working on this?

Since 2014 it's been just me, though I do plan to start a company and hire talented people to help me out with this project once the stars align.

## Can I show my support for this site? Do you accept donations?

I'm planning to add "pro accounts" for an extended feature set. You do not have to donate without getting something back, instead I'd rather be happy to see you profit from the donation as well. It would be my dream to work on this full-time.

## Can I help with coding or change stuff as this is Open Source?

Sure, the setup to start contributing is not that hard. Try to get in contact with me on Discord.

## Can I apply to be a data mod / editor?

Sure, just contact me on Discord if you want to help out with the database.