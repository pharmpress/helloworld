<?
$hashCode = crc32($_SERVER["SERVER_NAME"] . $_SERVER["SERVER_ADDR"] . $_SERVER["SERVER_PORT"]);
$mask = 0x1000000;
$bgcolor = $hashCode%$mask;
$fgcolor = ($mask - $bgcolor -1) %$mask;
?>
<html>
  <head>
    <title>Hello, World</title>
  </head>
  <body style="background-color:#<?= str_pad(dechex($bgcolor), 6, "0", STR_PAD_LEFT) ?>;color:#<?= str_pad(dechex($fgcolor), 6, "0", STR_PAD_LEFT) ?>">
	<h1>Hello, World!</h1>
    <p>
    	<ul>
    		<li>SERVER_NAME : <?= $_SERVER["SERVER_NAME"] ?></li>
    		<li>SERVER_ADDR : <?= $_SERVER["SERVER_ADDR"] ?></li>
    		<li>SERVER_PORT : <?= $_SERVER["SERVER_PORT"] ?></li>
    		<li>REMOTE_ADDR : <?= $_SERVER["REMOTE_ADDR"] ?></li>
    		<li>REMOTE_PORT : <?= $_SERVER["REMOTE_PORT"] ?></li>
 		<ul>
    </p>
  </body>
</html>