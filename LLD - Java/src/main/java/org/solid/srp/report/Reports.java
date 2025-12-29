package org.solid.srp.report;

class Report {
    String generateContent(){
        return "Context";
    }
}

class PDFPrinter{
    void print(Report report){

    }
}

class FileStorage{
    void save(String content){}
}
