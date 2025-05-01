const express = require("express");
const path = require("path");

const PORT = 5500;
const app = express();

app.use(express.static(path.join(__dirname,"public")));

app.get("/",(req,res)=>{
    res.sendFile(path.join(__dirname,"public","index.html"));
});

app.get("/create-playlist",(req,res)=>{
    res.sendFile(path.join(__dirname,"public","form.html"));
});

app.listen(PORT,()=>{
    console.log(`Frontend server is running on: ${PORT}`);
});