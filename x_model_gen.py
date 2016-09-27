"""
This is a small script for generating the initial Go model from the 
olca-schema yaml files. To run this script you need to have PyYAML 
installed:

    pip install pyyaml
    
You also have to configure the YAML_DIR in this script to point to
the directory where the YAML files are located: 

    # clone the olca-schema repository to some folder
    cd <folder>
    git clone https://github.com/GreenDelta/olca-schema.git
    # <folder>/olca-schema/yaml is the path for the YAML_DIR
    
After this you can run this script. It will print the generated structs
to the console:

    python x_model_gen.py > [.. path to generated file].go
    
"""

YAML_DIR = 'C:/Users/Besitzer/Downloads/olca-schema/yaml'


import yaml
from os import listdir


def print_class(class_model):
    name = class_model['name']
    print('// %s http://greendelta.github.io/olca-schema/html/%s.html' % (name, name))
    t = 'type %s struct {\n' % name
    if 'superClass' in class_model:
        t += '\t%s\n' % class_model['superClass']
    if 'properties' in class_model:
        for prop in class_model['properties']:
            t += '\t' + convert_property(prop) + '\n'
    t += '}\n'
    print(t)
    print_constructor(class_model)


def convert_property(prop):
    name = prop['name']    
    t = name[0].upper() + name[1:]
    type = prop['type']
    if type == 'integer':
        t += ' int' + (' `json:"%s"`' % name)
    elif type == 'double':
        t += ' float64' + (' `json:"%s"`' % name)
    elif type == 'boolean':
        t += ' bool' + (' `json:"%s"`' % name)
    elif type == 'date' or type == 'dateTime':
        t += ' string' + (' `json:"%s,omitempty"`' % name)
    elif type == 'List[string]':
        t += ' []string' + (' `json:"%s,omitempty"`' % name)
    elif type.startswith('List['):
        sub = type[5:(len(type)-1)]
        t += ' []' + sub + (' `json:"%s,omitempty"`' % name)    
    else:
        t += ' ' + type + (' `json:"%s,omitempty"`' % name)
    return t


def print_constructor(class_model):
    if 'superClass' not in class_model:
        return
    name = class_model['name']
    s = class_model['superClass']
    if s != 'RootEntity' and s != 'CategorizedEntity':
        return
    t = '// New%s initializes a new %s with the given id and name\n' % (name, name)
    v = name[0].lower()
    t += 'func New%s(id, name string) *%s {\n' % (name, name)
    t += '\t%s := %s{}\n' % (v, name)
    t += '\t%s.Context = ContextURL\n' % v
    t += '\t%s.Type = "%s"\n' % (v, name)
    t += '\t%s.ID = id\n' % v
    t += '\t%s.Name = name\n' % v
    t += '\treturn &%s\n' % v
    t += '}\n'    
    print(t)    


if __name__ == '__main__':
    print('package schema\n')
    for f in listdir(YAML_DIR):
        path = YAML_DIR + '/' + f
        with open(path, 'r', encoding='utf-8') as stream:
            model = yaml.load(stream)
            if 'class' in model:
                print_class(model['class'])
            