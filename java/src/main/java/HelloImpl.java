import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class HelloImpl implements HelloHandler {
    public static void main(String [] args) {
        new HelloServer(new HelloImpl()).run(4080);
    }

    public Greeting getGreeting(ResourceContext context, String name) {
        return new Greeting().message("Hello, " + name);
    }

    public ResourceContext newResourceContext(HttpServletRequest request, HttpServletResponse response) {
        return new HelloContext(request, response);
    }

    class HelloContext implements ResourceContext {
        HttpServletRequest request;
        HttpServletResponse response;
        HelloContext(HttpServletRequest request, HttpServletResponse response) {
            this.request = request;
            this.response = response;
        }
        public HttpServletRequest request() { return request; }
        public HttpServletResponse response() { return response; }
        public void authenticate() { }
        public void authorize(String action, String resource, String trustedDomain) { }
    }

}
