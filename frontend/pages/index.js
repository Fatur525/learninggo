import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import { useState, useEffect } from 'react'

export default function Home() {
  const [data, setData] = useState(null)
  const [error, setError] = useState(null)
  useEffect(() => {
    getData()
  }, []);

  console.log(data?.data)
  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL)
      const data = await res.json();
      setData(data)
    } 
    catch (error) {
      setError(error);
    }
  }

  return (
    <div>
      {error && <div>Failed to load (error.toString())</div>}
      {
        !data ? <div>Loading...</div>
          : (
            (data?.data ?? []).length === 0 && <p>data kosong</p>
          )
      }
      <Input onSuccess={getData}/>
      {data?.data && data?.data?.map((item, index) => (
        <p key={index}>{item}</p>
      ))}
    </div>
  
  )
}


function Input({onSuccess}) {
  const [data, setData] =  useState(null);
  const [error, setError] =  useState(null);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const body = {
      text: formData.get("data")
    }

    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`,{
        method :'POST',
        body : JSON.stringify(body)
      });
      const data = await res.json();
      setData(data.message);
      onSuccess();
    } 
    catch (error) {
      setError(error);
    }
  }
  return (
    <div>
      {error && <p>error: {error.toString()}</p>}
      {data && <p>success: {data}</p>}
      <form onSubmit={handleSubmit}>
        <input name="data" type="text" />
        <button>submit</button>
      </form>
    </div>
  )
}

