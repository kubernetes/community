python3 tm.py --dfd > updated-dataflow.dot
dot -Tpng < updated-dataflow.dot > updated-dataflow.png
open updated-dataflow.png
