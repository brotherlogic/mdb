kubectl delete secret ghb -n mdb
kubectl create secret generic ghb --from-literal ghb_password=$1 -n mdb