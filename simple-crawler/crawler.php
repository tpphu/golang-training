<?php

// @TODO
// 1. Download content
// 2. Extract content

function getTagContent($doc, $tag) {
  $start = strpos($doc, '<'.$tag.'>');
  $end = strpos($doc, '</'.$tag.'>');
  $title = substr($doc, $start + strlen($tag) + 2, $end - ($start + strlen($tag) + 2));
  return trim($title);
}

echo "\n\n\n";

function downloadAndExtract($url) {
  
  $body = file_get_contents($url);

  $title = getTagContent($body, 'title');

  echo $title ."\n";
}
function getUrlsFromDB() {
  return [
    'https://reqres.in/api/users?delay=3',
    'https://reqres.in/api/users?delay=2',
    'https://www.thesaigontimes.vn/274113/bao-giay-van-thu-vi.html',
  ];
}


$urls = getUrlsFromDB();
for($i = 0; $i < count($urls); $i++) {
  $url = $urls[$i];
  downloadAndExtract($url);
}
