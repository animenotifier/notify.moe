<?php
	$data = json_decode(apc_fetch('nyaa-errors'), true);
	unset($data['MAL is blocking ARN from accessing your list']);

	header('Content-Type: application/json');
	echo json_encode($data);
?>