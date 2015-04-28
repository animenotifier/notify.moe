<?php
	require_once("fw/fw.php");
?>
<!DOCTYPE html>
<html lang="en">
<head>
	<?php
		if($loggedIn) {
			$userName = $_SESSION['accountId'];;

			$pages['login']['visible'] = false;
			$pages['register']['visible'] = false;
			$pages['profile']['visible'] = true;
			$pages['profile']['url'] = "+$userName";
			$pages['feedback']['visible'] = true;
			//$pages['profile']['ajax'] = false;
		} else {
			$pages['logout']['visible'] = false;
		}
	?>
	<?php
		require_once("fw/header.php");
	?>

	<link rel="chrome-webstore-item" href="https://chrome.google.com/webstore/detail/hajchfikckiofgilinkpifobdbiajfch">
</head>

<body>
	<div id="container">
		<!-- Header -->
		<div id="header-container">
			<div id="header">
				<div id="title">
					<a href="/">
						<h1><?php echo $title; ?></h1>
					</a>
				</div>

				<div id="tagline">
					<!--Don't miss the next episode.--> <span style="opacity: 1.0;">See <a href="https://twitter.com/animenotifier" target="_blank">Twitter</a> for the latest news. KissAnime is currently not working because of their website changes.</span>
				</div>

				<nav id="navigation">
					<?php require_once("fw/navigation.php"); ?>
				</nav>
			</div>
		</div>
		
		<!-- Content -->
		<div id="content-container">
			<div id="content"><?php require_once("fw/content.php"); ?></div>
		</div>
		
		<!-- Footer -->
		<div id="footer-container">
			<footer id="footer">
				<div class="copyright">
					&copy; <?php echo date("Y") . " "; ?>
					
					<a href="http://blitzprog.org" target="_blank" itemscope itemtype='http://schema.org/Person'>
						<span itemprop="name">
							<?php echo $config["author"]; ?>
						</span>
					</a>
					|
					<a href="https://www.facebook.com/pages/Anime-Release-Notifier/1400941563536030" target="_blank" title="Facebook">
						Facebook
					</a>
					|
					<a href="https://twitter.com/animenotifier" target="_blank" title="Twitter">
						Twitter
					</a>
					|
					<a href="https://www.google.com/+AnimeReleaseNotifierOfficial" rel="publisher" target="_blank">
						Google+
					</a>
					|
					<a href="https://github.com/freezingwind/animereleasenotifier.com" target="_blank" title="GitHub">
						GitHub
					</a>
					|
					<!--<a href="/recruit" class="ajax">
						Recruiting data editors!
					</a>
					|-->
					<a href="/irc" class="ajax">
						IRC
					</a>
					|
					<a href="/roadmap" class="ajax">
						Roadmap
					</a>
					|
					<a href="/legal" class="ajax">
						Legal notice
					</a>
				</div>
			</footer>
		</div>
	</div>
	
	<!-- Loading animation -->
	<div id="loading-animation" class="spinner">
		<div class="bounce1"></div>
		<div class="bounce2"></div>
		<div class="bounce3"></div>
	</div>

	<!-- Scripts -->
	<?php
		include_once("fw/scripts.php");
		include("scripts/javascript.php");
	?>
</body>
</html>