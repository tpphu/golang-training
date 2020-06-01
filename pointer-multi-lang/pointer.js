// Struct Go
function Point(x, y) {
  this.x = x;
  this.y = y;
}
// Value P1 = [10, 11]
var p1 = new Point(10, 11);

console.log(p1);

// P2 = P1 
// ca p2 & p1 deu refer den cung 1 gia tri
var p2 = p1;
p2.x = -9;

console.log(p1)