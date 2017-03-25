
public class Hello {
    
    public static void main(String [] args) {
        String name = System.getenv("USER");
        Greeting greeting = new HelloClient("http://localhost:4080/hello/v1").getGreeting(name);
        System.out.println(greeting.message);
    }
}

