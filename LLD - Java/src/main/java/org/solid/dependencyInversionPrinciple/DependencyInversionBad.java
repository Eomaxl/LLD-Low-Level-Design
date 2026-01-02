package org.solid.dependencyInversionPrinciple;

public class DependencyInversionBad {
    
}

class EmailSender {
    void send(String message) {/* send mail */}
}

class NotificationService {
    EmailSender emailSender = new EmailSender();

    void notify(String message) {
        emailSender.send(message);
    }
}