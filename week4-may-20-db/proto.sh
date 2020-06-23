# docker run <image> ~ go run main.go
# docker run znly/protoc
# --rm		Automatically remove the container when it exits
# -v volume mapping cai $(pwd) thu muc hien tai ma minh dang dung, va tao mot folder tuong ung ben trong container, map 2 cai nay lai voi nhau
# => Y nghi cua no la: ben trong container se co cai folder tuong duong voi: /Users/tranphongphu/workspace/golang-training/week4-may-20-db/proto
# => Va no tro ra cai folder ben ngoai cua minh
# => Dieu do co nghia la, neu minh sua cai gi ben ngoai thi ben trong folder do cung thay doi, va nguoc lai
# -w: working directory, tuc la de khi minh chay mot cau lenh, no biet la chay cau lenh do o thu muc nao
# --go_out=. -I. proto/product.proto la cau lenh cua image nha
docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --go_out=plugins=grpc:. -I. proto/product.proto
# docker run --rm -v $(pwd):$(pwd) -w $(pwd) znly/protoc --gofast_out=plugins=grpc:. -I. proto/product.proto