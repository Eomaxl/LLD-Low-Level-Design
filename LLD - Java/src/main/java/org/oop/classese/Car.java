package org.oop.classese;

public class Car {
    String brand;
    String model;
    int speed;

    public Car(String brand, String model){
        this.brand = brand;
        this.model = model;
        this.speed = 0;
    }

    public void accelerate(int increment){
        this.speed += increment;
    }

    public void displayInfo(){
        System.out.println("Brand is "+this.brand+" , model is : "+this.model+" and speed is :"+this.speed);
    }

}
