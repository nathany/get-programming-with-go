public class Person {
   int age;

   public void birthday() {
      age++;
   }

   public static void main(String []args) {
      Person nobody = null;
      System.out.println(nobody);

      nobody.birthday();
   }
}
