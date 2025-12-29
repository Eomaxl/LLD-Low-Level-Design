class Report:
    def generate_content(self) -> str:
        return "content"
    
class PDFPrinter:
    def print_to_pdf(self, report: Report) ->None:
        # Instruction to print to pdf
        pass

class FileStorage:
    def save_to_file(self,content: str) -> None:
        # file I/O
        pass