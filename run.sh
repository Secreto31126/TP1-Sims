#How this script should be called:
#python3 config/config.py | xargs ./run.sh

M=$1
L=$2
N=$3
Rc=$4
Size=''
method=$5
loop=$6
particle=$7

#Remove old files
rm ./files/Static
rm ./files/Dynamic
rm ./output.txt

#Generate input for backend
go run ./generators/generate.go "$N" "$L" > ./files/Static.txt
go run ./generators/generate.go "$N" "$L" "d" > ./files/Dynamic.txt

#Run algorithm
go run ./main.go "$M" "$Rc" "" "$method" "$loop" > ./output.txt

#Plot results
python3 ./rendering/render.py "output.txt" "$M" "$L" "$particle"

