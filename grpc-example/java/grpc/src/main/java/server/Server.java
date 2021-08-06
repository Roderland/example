package server;

import com.google.protobuf.Empty;
import generate.Person;
import generate.PersonServiceGrpc;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * @author Roderland
 * @since 1.0
 */

public class Server {
    public static void main(String[] args) throws IOException {
        final io.grpc.Server server = ServerBuilder
                .forPort(50051)
                .addService(new PersonServiceImpl())
                .build()
                .start();

        while (true);
    }
}

class PersonServiceImpl extends PersonServiceGrpc.PersonServiceImplBase {

    private List<Person> dataSource;

    public PersonServiceImpl() {
        dataSource = new ArrayList<>();

        final Person.PhoneNumber.Builder phoneBuilder = Person.PhoneNumber
                .newBuilder()
                .setNumber("12345678910")
                .setType(Person.PhoneType.MOBILE);

        final Person.Builder personBuilder = Person
                .newBuilder()
                .setName("小明")
                .setId(1)
                .setEmail("996@icu.com")
                .addPhones(phoneBuilder.build());

        final Person person = personBuilder.build();

        dataSource.add(person);
    }

    @Override
    public void addPerson(Person request, StreamObserver<Person> responseObserver) {
        dataSource.add(request);
        responseObserver.onNext(request);
        responseObserver.onCompleted();
    }

    @Override
    public void listPerson(Empty request, StreamObserver<Person> responseObserver) {
        for (Person person : dataSource) {
            responseObserver.onNext(person);
        }
        responseObserver.onCompleted();
    }
}


