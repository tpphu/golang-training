function Point(x, y) {
  this.x = x;
  this.y = y;
}

var p1 = new Point(10, 11);

console.log(p1);

var p2 = p1;
p2.x = -9;

console.log(p1)