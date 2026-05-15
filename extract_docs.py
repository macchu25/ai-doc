import docx
import os
import sys

# Set output encoding to UTF-8
if sys.stdout.encoding != 'utf-8':
    sys.stdout.reconfigure(encoding='utf-8')

def extract_text(file_path):
    doc = docx.Document(file_path)
    full_text = []
    for para in doc.paragraphs:
        full_text.append(para.text)
    return '\n'.join(full_text)

files = [
    'task-1.1.1-user-requirements.docx',
    'task-1.1.2-use-case-analysis.docx',
    'task-1.1.3-output-catalog.docx'
]

for file in files:
    print(f"--- {file} ---")
    try:
        print(extract_text(file))
    except Exception as e:
        print(f"Error reading {file}: {e}")
    print("\n" + "="*50 + "\n")
