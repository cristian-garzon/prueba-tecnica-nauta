create table clients (
    id serial primary key,
    name varchar(255) not null,
    last_name varchar(255) not null,
    email varchar(255) not null unique,
    DNI bigint not null unique,
    created_at timestamp not null default now()
);

create table bookings (
    booking_id varchar(255) primary key,
    client_id bigint not null,
    status varchar(255) not null,
    origin_port varchar(255) not null,
    destination_port varchar(255) not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    foreign key (client_id) references clients(id)
);

create table containers (
    container_id varchar(255) primary key,
    booking_id varchar(255) not null,
    container_type varchar(255) not null,
    description varchar(255) not null,
    weight float not null,
    created_at timestamp not null default now(),
    foreign key (booking_id) references bookings(booking_id)
);

create table orders (
    purchase_id varchar(255) primary key,
    booking_id varchar(255) not null,
    status varchar(255) not null,
    total_amount bigint not null,
    description varchar(255) not null,
    created_at timestamp not null default now(),
    foreign key (booking_id) references bookings(booking_id)
);

create table invoices (
    invoice_id varchar(255) primary key,
    purchase_id varchar(255) not null,
    amount bigint not null,
    status varchar(255) not null,
    payment_date timestamp not null,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    foreign key (purchase_id) references orders(purchase_id)
);