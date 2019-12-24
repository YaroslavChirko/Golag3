let fs = require('fs');
// var buf = Buffer.from('./test.html');
  // fs.createReadStream(buf).pipe(res);
  // let buffer = new Buffer.alloc(10);
 let Chunk = 10*1024*1024,
     buffer = Buffer.alloc(Chunk);

let pathR = process.argv[2],
    pathW = process.argv[3];

fs.exists(pathW, exists => {
   if (!exists) fs.mkdir(pathW, err => {if (err) throw err});
});  
     
fs.readdir(pathR, (err, files)=>{
	 if (err) throw err
	 let counter = 0;
     files.forEach(function(file){
     	
     	// console.log(file.toString());
     	if(file.slice(-4) === ".txt"){
     		counter++; 
      
   //  fs.readFile(pathR + '/' + file, 'utf8', (err, data)=>{    	
  fs.open(pathR + '/' + file, 'r', function(err, fd){
        if (err) throw err;
   function readNextChunk() {
    fs.read(fd, buffer, 0, Chunk, null, function(err, nread) {
      if (err) throw err;
         if (nread === 0) {
       

        fs.close(fd, function(err) {
          if (err) throw err;
        });
        return;
      }

      let data;
      if (nread < Chunk)
        data = buffer.slice(0, nread);
      else
        data = buffer;
 //console.log(data.toString());
   
        fs.writeFile(pathW + '/' + file.split('.')[0] + ".res",
        fixString(data.toString()),
        err => {
    	    if (err) throw err
        })  
    
   });
} readNextChunk();  //readChunk
})
     }});
 console.log("Total number of processed files " +counter);
});    
     
function fixString(inp){
    let i, len, arr, outp, Rside, Lside, RsideIsNum, LsideIsNum;
    arr = inp.split(",");
    outp = "";
    for(i=1, len=arr.length; i<len; i++){
        Lside = arr[i-1];
        Rside = arr[i];
        LsideIsNum = /[a-z]|[A-Z]/.test(Lside.charAt(Lside.length-1));
        RsideIsNum = /[a-z]|[A-Z]/.test(Rside.charAt(0));
        outp +=  ","+((LsideIsNum || RsideIsNum)?" ":"") + Rside;
    }
    return (arr[0] + outp).replace(/\s\s+/g," ");
}     
    