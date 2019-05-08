<?php

class Point {
  var $x = 0;
  var $y = 0;

  public function __construct($x, $y){
    $this->x = $x;
    $this->y = $y;
  }
}

$p1 = new Point(10, 11);

var_dump($p1);

$p2 = $p1;
$p2->x = -9;

var_dump($p1);