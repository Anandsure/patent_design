import os
from zipfile import ZipFile
import traceback

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


