package client;

import com.google.protobuf.Empty;
import generate.Person;
import generate.PersonServiceGrpc;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import java.util.Iterator;

/**
 * @author Roderland
 * @since 1.0
 */
public class Client {
    PersonServiceGrpc.PersonServiceBlockingStub blockingStub;

    public Client() {
        final ManagedChannel channel = ManagedChannelBuilder
                .forAddress("localhost", 50051)
                .usePlaintext()
                .build();
        blockingStub = PersonServiceGrpc.newBlockingStub(channel);
    }

    public static void main(String[] args) {

        final Client client = new Client();

        final Person.PhoneNumber phoneNumber = Person.PhoneNumber.newBuilder()
                .setNumber("12345678900")
                .setType(Person.PhoneType.HOME)
                .build();

        final Person person = Person.newBuilder()
                .setId(2)
                .setName("华强")
                .setEmail("007@icu.com")
                .addPhones(phoneNumber)
                .build();

        final Person addPerson = client.blockingStub.addPerson(person);
        System.out.println(addPerson);

        final Iterator<Person> listPerson = client.blockingStub.listPerson(Empty.newBuilder().build());
        listPerson.forEachRemaining(System.out::println);
    }
}
