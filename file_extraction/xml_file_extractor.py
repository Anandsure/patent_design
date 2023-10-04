import os
from zipfile import ZipFile
import traceback



def extract_all_files():
    directory = "datasets/I20230502/I20230502/DESIGN/"
    export_path = "file_extraction/full_data/"
    for filename in os.listdir(directory):
        try:
            f = os.path.join(directory, filename)
            with ZipFile(f, 'r') as zObject:
                zObject.extractall(path=export_path)
            zObject.close()
        except Exception as err:
            print(err)
            traceback.print_exc()

def move_all_xml_files():
    export_path = "file_extraction/all_xml/"
    directory = "file_extraction/full_data/"
    for filename in os.listdir(directory):
        try:
            f = os.path.join(directory, filename)
            xml_file = f+"/"+filename+".XML"
            os.replace(xml_file,export_path+filename+".XML")
        except Exception as err:
            print(err)
            traceback.print_exc()

move_all_xml_files()