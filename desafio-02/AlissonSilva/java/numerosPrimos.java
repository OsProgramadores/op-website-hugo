public class numerosPrimos {
    public static void main(String[] args) {

        int num = 0;

        for (int i = 0; i < 10000; i++) {
            if (numPrimo(i)) {
                System.out.println(i + " é primo.");
            }
        }
    }

    private static boolean numPrimo(int numero) {

        for (int j = 2; j < numero; j++) {
            if (numero % j == 0) {
                
                return false;
            }
        }

        return true;
    }
}