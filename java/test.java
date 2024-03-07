public class test {
    private static int t;
    public static void main(String[] args) {
        String s;
        s = String.join(",", args);
        System.out.println(s);
        t = 10;
        test(6);
    }

    private static void test(int i) {
        System.out.println(i);
        System.out.println(t);
    }
}



