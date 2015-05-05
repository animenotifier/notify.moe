<?php
	// Setup
	set_include_path($_SERVER['DOCUMENT_ROOT']);
	require_once('php-paypal-ipn/IPNListener.php');
	require_once('fw/helper.php'); // For sending slack messages
	require_once('config.php'); // For aerospike config

	ini_set('log_errors', true);
	//ini_set('error_log', dirname(__FILE__) . DIRECTORY_SEPARATOR .  'ipn_errors.log');

	// Aerospike
	$db = new Aerospike($config["aeroSpike"]);

	// PayPal
	$paypal = new IPNListener();
	$paypal->use_sandbox = true;

	// Config
	$mailAddress = 'e.urbach@gmail.com';

	if($verified = $paypal->processIpn()) {
		$transactionRawData = $paypal->getRawPostData();
		$transactionData = $paypal->getPostData();

		// 1. Check that $_POST['payment_status'] is "Completed"
		if($_POST['payment_status'] != 'Completed') {
			sendSlackMessage('Denied (payment status not completed):\n```' . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		// 2. Check that $_POST['txn_id'] has not been previously processed
		$transactionKey = $db->initKey('arn', 'Transactions', $_POST['txn_id']);
		$status = $db->exists($transactionKey, $metadata);

		if($status == Aerospike::OK) {
			sendSlackMessage("Denied (transaction ID has already been processed):\n```" . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		// 3. Check that $_POST['receiver_email'] is your Primary PayPal email
		if($_POST['receiver_email'] !== $mailAddress) {
			sendSlackMessage("Denied (wrong mail address):\n```" . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		// 4. Check that $_POST['mc_gross'] is correct
		if(floatval($_POST['mc_gross']) <= 0) {
			sendSlackMessage("Denied (payment amount is 0 or lower than 0):\n```" . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		// 5. Check that $_POST['mc_currency'] is correct
		if($_POST['mc_currency'] !== 'USD') {
			sendSlackMessage("Denied (wrong currency):\n```" . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		$userName = $_POST['custom'];

		// Transaction data
		$status = $db->put($transactionKey, [
			'id' => $_POST['txn_id'],
			'amount' => $_POST['mc_gross'],
			'currency' => $_POST['mc_currency'],
			'email' => $_POST['payer_email'],
			'firstName' => $_POST['first_name'],
			'lastName' => $_POST['last_name']
		]);

		// Store transaction in database
		if($status !== Aerospike::OK) {
			sendSlackMessage("Error storing transaction in database:\n```" . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		// Get account
		$userKey = $db->initKey('arn', 'Users', $userName);
		$status = $db->get($userKey, $record);

		// Error retrieving account?
		if($status !== Aerospike::OK) {
			sendSlackMessage("Error retrieving account: $userName\n```" . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		$user = $record['bins'];
		if(array_key_exists('balance', $user)) {
			$user['balance'] = floatval($user['balance']) + floatval($_POST['mc_gross']);
		} else {
			$user['balance'] = floatval($_POST['mc_gross']);
		}

		// Save back to database
		$status = $db->put($userKey, $user);

		// Error saving account?
		if($status !== Aerospike::OK) {
			sendSlackMessage("Error saving account: $userName\n```" . $paypal->getTextReport() . '```', 'paypal-errors');
			exit;
		}

		sendSlackMessage("Success!\n```" . $paypal->getTextReport() . '```', 'paypal-success');
		sendSlackMessage("$userName | https://animereleasenotifier.com/+$userName | " . $_POST['payer_email'] . ' | +' . $_POST['mc_gross'] . ' ' . $_POST['mc_currency'] . ' => ' . $user['balance'], 'payments');
	} else {
		$errors = $paypal->getErrors();
		sendSlackMessage("Not verified:\n```" . print_r($errors, true) . print_r($_POST, true) . '```', 'paypal-errors');
	}
?>