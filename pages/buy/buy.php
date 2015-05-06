<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('fw/login.php');
	require_once('config.php');

	if(!$loggedIn) {
		echo "Not logged in!";
		return;
	}

	$userName = $_SESSION['accountId'];

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
	$balance = intval($user['balance'] * 100);
?>

<style scoped>
	.balance {
		float: left;
		width: 100%;
		text-align: right;
		font-size: 2em;
		font-weight: bold;
	}

	.balance:after {
		content: "💰";
	}
</style>

<p>
	<span class="balance" title="You currently own <?php echo $balance; ?> credits.">
		<?php echo $balance; ?>
	</span>
</p>

<h2>Get credits</h2>
<p>
	<a href="https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=MSTSA6EXDH2GY&item_name=200 credits (for <?php echo $userName; ?>)&custom=<?php echo $userName; ?>" target="_blank">200 💰</a>
</p>

<h2>BASIC subscription</h2>
<p>
	<a><img src="https://assets-cdn.github.com/images/icons/emoji/unicode/1f4e6.png" alt="package"> 1 Year (199 credits)</a>
</p>

<h2>PRO subscription</h2>
<p>
	<a><img src="https://assets-cdn.github.com/images/icons/emoji/unicode/1f4e6.png" alt="package"> 1 Year (1999 credits)</a>
</p>