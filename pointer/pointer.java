class Point {
  int x = 0;
  int y = 0;
  public Point(int x, int y) {
    this.x = x;
    this.y = y;
  }
}
class Main {
  public static void main(String[] args) {
    Point p1 = new Point(10, 11);
    System.out.println("[x = "+p1.x +",y="+ p1.y+"]");
    Point p2 = p1;
    p2.x = -9;
    System.out.println("[x = "+p1.x +",y="+ p1.y+"]");
  }
}