<?php
interface ListProvider {
	public function getAnimeList($userName, $completed = false);
	public function getAnimeListUrl($userName);
	public function clearCache($userName);
}
?>