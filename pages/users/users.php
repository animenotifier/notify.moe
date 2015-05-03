<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/page.php');
	require_once('fw/login.php');
	
	if(count($params) == 0) {
		//$loginSystem->showAllUsers();
		return;
	}
	
	$userName = $params[0];

	if($userName == '') {
		return;
	}
?>

<style scoped>
	tr {
		transition: background 270ms ease;
	}

	tr td:first-child {
		border-top-left-radius: 3px;
		border-bottom-left-radius: 3px;
	}

	tr td:last-child {
		border-top-right-radius: 3px;
		border-bottom-right-radius: 3px;
	}

	tr:hover {
		/*background: rgba(0, 0, 0, 0.2);*/
	}

	td {
		font-size: 1.1em;
		vertical-align: middle;
	}

	.label {
		width: 20%;
		height: calc(3.7em - 2em);
		padding: 1em;
	}

	.profile-image {
		border-radius: 5px;
		box-shadow: 0 0 8px rgba(0, 0, 0, 0.5);
		margin-right: 1em;
	}

	.profile-image-container {
		float: left;
		text-align: left;
		width: auto;
	}

	.user-info {
		float: left;
		width: auto !important;
		margin-left: 1.5em;
	}

	.user-name {
		float: left;
		text-align: left;
		display: block;
		width: 100%;

		font-size: 2.8em;
		line-height: 1em;
		
		margin: 0.25em 0;
		letter-spacing: 0px;
	}

	.user-name a:hover {
		text-shadow: none;
	}

	.user-title {
		float: left;
		text-align: left;
		display: block;
		width: 100%;

		font-weight: normal;
		font-size: 0.9em;
		opacity: 0.5;
		margin-top: 1em 0;
		line-height: 1em;
		letter-spacing: 1px;
	}

	.user-tagline,
	.user-website {
		float: left;
		text-align: left;
		display: block;
		width: auto;
		opacity: 0.7;
		margin-top: 1em;
	}

	.user-website {
		margin-left: 0.6em;
		opacity: 1;
	}

	input[type="text"] {
		max-width: 250px;
	}

	.wip {
		opacity: 0.2;
	}

	@media only screen and (max-width: 800px) {
		table {
			display: block;
		}

		td {
			display: block !important;
			border-bottom: none !important;
			vertical-align: middle !important;
		}

		.label {
			width: 100%;
			height: auto;
		}

		table {
			max-width: 100%;
		}
	}
</style>

<?php
	$db = new Aerospike($config["aeroSpike"]);

	if(!$db->isConnected()) {
		echo "Failed to connect to the database [{$db->errorno()}]: {$db->error()}\n";
		exit(1);
	}

	$key = $db->initKey("arn", "Users", $userName);
	$status = $db->get($key, $record);

	if($status == Aerospike::ERR_RECORD_NOT_FOUND) {
		echo "The user ". $key['key']. " does not exist in the database.\n";
		exit(2);
	}

	if($status != Aerospike::OK) {
		echo $db->error();
		exit(3);
	}

	$user = $record['bins'];
	$providers = $user['providers'];
	$listProviderName = $providers['list'];
	$animeProviderName = $providers['anime'];
	$timeProviderName = $providers['time'];
	$email = $user['email'];
	$sortBy = $user['sortBy'];

	$accountKey = $db->initKey('arn', 'Accounts', $email);
	$db->get($accountKey, $accountRecord);
	$account = $accountRecord['bins'];

	$db->close();

	// Own account
	$disabled = '';
	$ownAccount = true;

	if(!$loggedIn || $_SESSION['accountId'] != $userName) {
		$disabled = 'disabled';
		$ownAccount = false;
	}
?>
<?php
	if($account['activated'] === 0) {
		echo "<div class='error'>Account hasn't been activated yet.</div>";

		if($ownAccount) {
			echo "
				<div>
					If your activation mail is not in the spam folder you can <a href='/resendmail' class='ajax'>resend the activation mail</a>.
				</div>";
		}
	}

	// List
	function listProviderOption($name, $value) {
		global $listProviderName;

		$selected = '';

		if($value == $listProviderName)
			$selected = " selected";

		echo "<option value='$value'$selected>$name</option>";
	}

	// Anime
	function animeProviderOption($name, $value) {
		global $animeProviderName;

		$selected = '';

		if($value == $animeProviderName)
			$selected = " selected";

		echo "<option value='$value'$selected>$name</option>";
	}

	// Time
	function timeProviderOption($name, $value) {
		global $timeProviderName;

		$selected = '';

		if($value == $timeProviderName)
			$selected = " selected";

		echo "<option value='$value'$selected>$name</option>";
	}

	// Sort
	function sortOption($name, $value) {
		global $sortBy;

		$selected = '';

		if($value == $sortBy)
			$selected = " selected";

		echo "<option value='$value'$selected>$name</option>";
	}

	$gravatarHash = md5(strtolower($email));
?>

<p class="profile-image-container">
	<?php
		if($ownAccount)
			echo '<a href="https://gravatar.com/" title="Change your globally recognized avatar" target="_blank">';
	?>

	<img src='https://www.gravatar.com/avatar/<?php echo $gravatarHash; ?>?s=200' alt='Avatar' class='profile-image'>
	
	<?php
		if($ownAccount)
			echo '</a>';
	?>

	<div class="user-info">
		<span class="user-name"><a href="/+<?php echo $userName; ?>" class="ajax"><?php echo $userName; ?></a></span>

		<?php
			$userTitle = "";

			$admins = array('Aky');
			$dataMods = array('mrskizzex', 'Bleack', 'Jinnial');
			$donators = array('Soulcramer', 'synthtech', 'drill', 'izzy', 'Josh', 'mrskizzex', 'magister');

			if(in_array($userName, $admins)) {
				$userTitle = "Admin";
			} else if(in_array($userName, $dataMods)) {
				$userTitle = "Data Editor";
			} else if(in_array($userName, $donators)) {
				$userTitle = "Honorable Member";
			} else {
				$userTitle = "Member";
			}

			if($userName == "Josh")
				$userTitle = "V.I.P. (admin of anilist.co)";

			if($userTitle !="") {
				echo "<span class='user-title'>$userTitle</span>";
			}
		?>

		<?php
			// Tagline
			if(array_key_exists('tagLine', $user))
				$tagLine = $user['tagLine'];
			else
				$tagLine = '';

			if($tagLine != '') {
				echo "<span class='user-tagline'>$tagLine</span>";
			} else {
				echo "<span class='user-tagline'><em>No tagline yet.</em></span>";
			}
		?>

		<?php
			// Website
			if(array_key_exists('website', $user))
				$website = $user['website'];
			else
				$website = '';

			if(preg_match("#https?://#", $website) === 0) {
				$url = "http://$website";
			} else {
				$url = $website;
			}

			echo "<a href='$url' target='_blank' class='user-website' rel='nofollow'>$website</a>";
		?>

		<?php
			/*if($animeProviderName === "Nyaa" && ($listProviderName === "AniList" || $listProviderName === "HummingBird"))
				echo "<span class='so-cool' style='float: left; width: 100%; text-align: left; margin-top: 1em; color: rgb(32, 255, 32)'>✓ Cool™</span>";*/
		?>
	</div>
</p>

<table>
	<tr>
		<td class="label">
			<img class="icon" itemprop="image" src="/images/icons/anime-list.png" alt="List provider" width="16" height="16">&nbsp; List:
		</td>
		<td>
			<select id="listProvider" <?php echo $disabled; ?>>
				<?php
					listProviderOption("anilist.co", "AniList");
					listProviderOption("anime-planet.com", "AnimePlanet");
					listProviderOption("hummingbird.me", "HummingBird");
					listProviderOption("myanimelist.net", "MyAnimeList");
				?>
			</select>
		</td>
	</tr>

<?php if($ownAccount): ?>
	<tr>
		<td class="label">
			<img class="icon" itemprop="image" src="/images/icons/list-username.png" alt="List user name" width="16" height="16">&nbsp; Name on <span class="list-provider-name-highlight"><?php echo $listProviderName; ?></span>:
		</td>
		<td>
			<?php
				$animeLists = $user['animeLists'];

				if(array_key_exists($listProviderName, $animeLists))
					$listUserName = $animeLists[$listProviderName]['userName'];
				else
					$listUserName = '';
			?>
			<input type="text" id="listUserName" placeholder="<?php echo $listProviderName; ?> user name" style="width: 220px;" value="<?php echo $listUserName; ?>" <?php echo $disabled; ?>>
		</td>
	</tr>
<?php endif; ?>

	<!-- Auth -->
	<?php if($ownAccount && $listProviderName == 'AniList'): ?>
	<tr style="display: none;">
		<td class="label">
			AniList PIN:
		</td>
		<td>
			<?php
				$aniList = $user['animeLists']['AniList'];

				if(array_key_exists('auth', $aniList))
					$auth = $aniList['auth'];
				else
					$auth = '';
			?>
			<input type="text" id="auth" placeholder="Needed if you want to edit the list" style="max-width: 450px;" value="<?php echo $auth; ?>">
			<br>
			<a href="https://anilist.co/api/auth/authorize?grant_type=authorization_pin&response_type=pin&client_id=akyoto-wbdln" target="_blank">Need a new one?</a>
			<br>
			<span style="opacity: 0.2">(work in progress)</span>
		</td>
	</tr>
	<?php endif; ?>

	<tr>
		<td class="label">
			<img class="icon" itemprop="image" src="/images/icons/anime.png" alt="Airing date" width="16" height="16">&nbsp; Anime:
		</td>
		<td>
			<select id="animeProvider" <?php echo $disabled; ?>>
				<?php
					animeProviderOption("nyaa.se", "Nyaa");
					animeProviderOption("animeshow.tv (ALPHA)", "AnimeShow");
					animeProviderOption("kissanime.com (deprecated)", "KissAnime");
					animeProviderOption("twist.moe (deprecated)", "AnimeTwist");
				?>
			</select>
		</td>
	</tr>

	<tr>
		<td class="label">
			<img class="icon" itemprop="image" src="/images/icons/airing-date.png" alt="Airing date" width="16" height="16">&nbsp; Airing time:
		</td>
		<td>
			<select id="timeProvider" <?php echo $disabled; ?>>
				<?php
					timeProviderOption("anilist.co", "AniList");
				?>
			</select>
		</td>
	</tr>

<?php if($ownAccount): ?>
	<tr>
		<td class="label">
			Tagline:
		</td>
		<td>
			<input type="text" id="tagLine" placeholder="Say something about yourself" maxlength="100" style="max-width: 220px;" value="<?php echo $tagLine; ?>">
		</td>
	</tr>

	<tr>
		<td class="label">
			Website:
		</td>
		<td>
			<input type="text" id="website" placeholder="Website URL" style="max-width: 220px;" value="<?php echo $website; ?>">
		</td>
	</tr>

	<tr>
		<td class="label">
			Sort by:
		</td>
		<td>
			<select id="sortBy" <?php echo $disabled; ?>>
				<?php
					sortOption("Airing date", "airingDate");
					sortOption("Alphabetically", "title");
				?>
			</select>
		</td>
	</tr>

	<tr>
		<td class="label">
			Opacity by:
		</td>
		<td class="wip">
			(work in progress)
		</td>
	</tr>

	<tr>
		<td class="label">
			Theme:
		</td>
		<td class="wip">
			(work in progress)
		</td>
	</tr>
<?php endif; ?>

	<tr>
		<td class="label" style="vertical-align: top;">
			Preview:
		</td>
		<td id="animeList">
			&nbsp;
		</td>
	</tr>

<?php if($ownAccount): ?>
	<tr>
		<td class="label" style="vertical-align: top;">
			Custom filters:
		</td>
		<td id="custom-filters">
			(work in progress)
		</td>
	</tr>
<?php endif; ?>

	<tr>
		<td class="label">
			Success rate:
		</td>
		<td id="successRate">
			-
		</td>
	</tr>

	<tr>
		<td class="label">
			Link:
		</td>
		<td>
			<a href="" target="_blank" id="listUrl"></a>
		</td>
	</tr>

	<tr>
		<td class="label">
			Export:
		</td>
		<td>
			<a href="http://animereleasenotifier.com/api/users/<?php echo $userName; ?>" target="_blank">
				User JSON
			</a>
			|
			<a href="http://animereleasenotifier.com/api/animelist/<?php echo $userName; ?>" target="_blank">
				Animelist JSON
			</a>
		</td>
	</tr>

	<!--<tr>
		<td class="label" style="vertical-align: top;">
			Comments:
		</td>
		<td>
			<div id="disqus_thread"></div>
		</td>
	</tr>-->
</table>

<script>
	document.addEventListener('DOMContentLoaded', setup);

	<?php require_once("js/anime-list.js"); ?>

	// Setup
	function setup() {
		<?php
		if($ownAccount) {
			echo "localStorage.userName = '$userName';";
		}
		?>

		setPageHandler(function(pageId) {
			document.removeEventListener('DOMContentLoaded', setup);
		});

		var saveSettings = function() {
			save(function() {
				document.removeEventListener('DOMContentLoaded', setup);

				$.get("/pages/users/users.php?params=<?php echo $userName; ?>", function(html) {
					$loadingAnimation.stop().fadeOut(fadeSpeed);
					$("#content").html(html);
					fireContentLoadedEvent();
				});
			});
		};

		$("select").change(saveSettings);
		$("input").change(function() {
			$(this).blur(saveSettings);
		});

		var $element = $("#animeList");
		var loading = "<div class='loading'><div class='rect1'></div><div class='rect2'></div><div class='rect3'></div><div class='rect4'></div><div class='rect5'></div></div>";
		
		$element.html(loading);

		var userName = "<?php echo $userName; ?>";
		var $customFilters = $("#custom-filters");
		var $listUrl = $("#listUrl");

		$.getJSON("/api/animelist/" + userName, function(json) {
			var animeList = new AnimeList(json, $element, 1 <?php if($ownAccount) echo ", function(anime){ anime.sendNotification(); }"; ?>);

			$("#successRate").text(Math.round(animeList.successRate * 100) + " %");

			$listUrl.attr("href", animeList.listUrl);
			$listUrl.text(animeList.listUrl);

			if(userName != "Aky")
				return;

			$customFilters.html("");

			json.watching.forEach(function(anime) {
				var titleInput = "<input type='text' placeholder='" + anime.title + "' title='If the title does not match with your anime provider enter the correct title here'>";
				var subsInput = "<select></select>";
				var qualityInput = "<select></select>";

				$customFilters.append("<div class='custom-filters-anime'>" + titleInput + "</div>");
			}.bind(this));
		}).fail(function() {
			$element.html("-");
		});
	}

	// Save
	function save(callBack) {
		var customFiltersString = "";

		var postData = {
			listProvider: $("#listProvider").val(),
			animeProvider: $("#animeProvider").val(),
			timeProvider: $("#timeProvider").val(),
			oldListProvider: "<?php echo $listProviderName; ?>",
			oldListUserName: $("#listUserName").val(),
			sortBy: $("#sortBy").val(),
			tagLine: $("#tagLine").val(),
			website: $("#website").val(),
			auth: "", //$("#auth").val(),
			customFilters: customFiltersString
		};

		$loadingAnimation.fadeIn(fadeSpeed);
		
		$.post("/pages/users/save-settings.php", postData, function(response) {
			//console.log(url);
			//console.log(response);

			callBack();
		});
	}
</script>