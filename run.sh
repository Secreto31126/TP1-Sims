#How this script should be called:
#python3 config/config.py | xargs ./run.sh

M=$1
L=$2
N=$3
Rc=$4
method=$5
particle=$6
Size=''
periodic=$7


#Remove old files
rm ./files/Static.txt
rm ./files/Dynamic.txt
rm ./output.txt

#Generate input for backend
go run ./generators/generate.go "$N" "$L" > ./files/Static.txt
go run ./generators/generate.go "$N" "$L" "d" > ./files/Dynamic.txt

#Run algorithm
go run ./main.go "$M" "$Rc" "" "$method" "$periodic" > ./output.txt

#Plot results
python3 ./rendering/render.py "output.txt" "$M" "$L" "$particle"

