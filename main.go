package main

import "io/ioutil"
import "io"
import "os"
import "fmt"
import "strings"
import "sync"
import "log"
import "path/filepath"


func remov( files [] os.FileInfo)[] os.FileInfo{
for j:=0; j<len(files); j++{
    	if filepath.Ext(files[j].Name()) != ".txt" {
    		files[j]=files[len(files)-1] 
    			j--
		files=files[:len(files)-1]
		}
		
    }
	return files
}

func output(file *os.File, err error, num int, name string , thisSes bool){
	checkFirst := false
	const BufferSize = 64
if err != nil {
  fmt.Println(err)
  return
}
defer file.Close()

buffer := make([]byte, BufferSize)

	for {
  		bytesread, err := file.Read(buffer)

  		if err != nil {
    		if err != io.EOF {
     		 fmt.Println(err)
    		}
    	break
  		}

  		if(bytesread<BufferSize){
  			buffer = buffer[:bytesread]
  		}
  		for i :=0;i<len(buffer);i++{
  				if(checkFirst == true&&i==0&&buffer[i]!=32){
  					siz := len(buffer)+1
    				buffertmp := make([]byte, siz)
    				buffertmp[0]=byte(32)
    				copy(buffertmp[1:],buffer[0:])
    				buffer=append(buffer,32)
    				buffer=buffertmp[:len(buffertmp)]
    				checkFirst = false
  				}else if(i < len(buffer)-1 && buffer[i]==44 && ((buffer[i+1]>=65&&buffer[i+1]<=90)||(buffer[i+1]>=97&&buffer[i+1]<=122))){

    				siz := len(buffer)+1
    				buffertmp := make([]byte, siz)
    				n := i+1
    				copy(buffertmp[:n],buffer[:n])
    				buffertmp[n]=byte(32)
    				copy(buffertmp[n+1:],buffer[n:])
    				buffer=append(buffer,32)
    				buffer=buffertmp[:len(buffertmp)]
    			}else if(i>0&&buffer[i]==44 &&((buffer[i-1]>=65&&buffer[i-1]<=90)||(buffer[i-1]>=97&&buffer[i-1]<=122))){
    					siz := len(buffer)+1
    				buffertmp := make([]byte, siz)
    				n := i+1
    				copy(buffertmp[:n],buffer[:n])
    				buffertmp[n]=byte(32)
    				copy(buffertmp[n+1:],buffer[n:])
    				buffer=append(buffer,32)
    				buffer=buffertmp[:len(buffertmp)]
    				}else if(i == len(buffer)-1 && buffer[i]==44){
    				checkFirst = true
    			}else{}
    	}

  		//fmt.Println("bytes read: ", bytesread)

  		filename := fmt.Sprintf("%s%s",name,".res")

  		if _, err := os.Stat("./"+os.Args[2]); os.IsNotExist(err) {
    os.Mkdir("./"+os.Args[2], os.ModeDir)
}		
	if thisSes == true{
			f, err := os.OpenFile("./"+os.Args[2]+"/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    		if err != nil {
        		log.Fatal(err)
    		}
    		_, err = f.Write(buffer)
    		if err != nil {
        		log.Fatal(err)
    		}

    		f.Close()
		}else{
			os.Remove("./"+os.Args[2]+"/"+filename)
			f, err := os.OpenFile("./"+os.Args[2]+"/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
    		if err != nil {
        		log.Fatal(err)
    		}
    		_, err = f.Write(buffer)
    		if err != nil {
        		log.Fatal(err)
    		}

    		f.Close()
    		thisSes=true
		}
  	




  		//fmt.Println("bytestream to string: ", string(buffer[:bytesread]))
	}
	fmt.Printf("File  %d \n",num)
}

func main() {
    files,_ := ioutil.ReadDir("./"+os.Args[1])
    
   
    files = remov(files)
    var wg sync.WaitGroup
    wg.Add(len(files))
    fmt.Printf("Num of files: %d \n",len(files))
 	
    for i := 0; i< len(files); i++ {
    	file,err := os.Open(os.Args[1]+"/"+files[i].Name())
    	go func (i int){
    		output(file,err, i,strings.Trim(files[i].Name(),".txt"),false)
    		wg.Done()
    		}(i)
    	
    }

    //fmt.Scanln();
    wg.Wait()
    fmt.Printf("Files affected: %d",len(files));


}