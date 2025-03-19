<?php
//include the quotes array
include 'quotes.php';

//pick a random quote
$random_quote = $quotes[array_rand($quotes)];
?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Random Quote Generator</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            text-align: center;
            margin: 50px;
        }
        .quote-box {
            padding: 20px;
            border: 2px solid #333;
            display: inline-block;
            font-size: 20px;
            max-width: 500px;
            background: #f9f9f9;
        }
        button {
            margin-top: 20px;
            padding: 10px 20px;
            font-size: 16px;
            cursor: pointer;
        }
        </style>
</head>
<body>
    <h1>Random Quote Generator</h1>
    <div class="quote-box">
       <p><?php echo $random_quote; ?></p> 
    </div>
    <br>
    <button onclick="window.location.reload();">Get Another Quote</button>
</body>
</html>