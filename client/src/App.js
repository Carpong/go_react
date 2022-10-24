import axios from 'axios'
import { useState } from 'react'

function App() {

  const [name, setName] = useState("");

  const [employeeList, setemployeeList] = useState([]);

  const getEmployee = ()=>{
    axios.get('http://localhost:8080/users').then(function (response) {
      setemployeeList(response.data.data)
    })
  }

  const addEmployee = ()=>{
    axios.post('http://localhost:8080/adduser',{
      name: name,
    }).then(function (response) {
      console.log(response);
    }).catch(function (error) {
      console.log(error);
    });
  }

  return (
    <div className="App container">
      <h1>Employee Information</h1>
      <div className="information">
        <form action="">
          <div className="mb-3" >
            <label htmlFor="name" className="form-label">Name :</label>
            <input type="text" className="form-control" placeholder="Enter Name" onChange={(event)=>{setName(event.target.val)}}></input>
          </div>
          {/* <div className="mb-3">
            <label htmlFor="age" className="form-label">Age :</label>
            <input type="number" className="form-control" placeholder="Enter Age"></input>
          </div>
          <div className="mb-3">
            <label htmlFor="mail" className="form-label">E-mail :</label>
            <input type="mail" className="form-control" placeholder="Enter E-mail"></input>
          </div> */}
          <button className="btn btn-success" onClick={addEmployee}>Add Employee</button>
        </form>
      </div>
      <hr/>
      <div className="employees">
        <button className="btn btn-primary" onClick={getEmployee}>Show Employee</button>
        <br/><br/>
        {employeeList.map((val, key)=>{
          return(
            <div className='employee card' key={key}>
              <div className='card-body text-left'>
                <p className='card-text' >Name : {val.name}</p>
              </div>
              <div className='d-flex'>
                <input type="text" placeholder="Enter Name" style={{width:"500px"}} className='form-control' ></input> 
                <button className='btn btn-warning'>update</button>
              </div>
            </div>
          )
        })}
      </div>
    </div>
  );
}

export default App;
