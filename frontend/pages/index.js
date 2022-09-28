import Head from 'next/head'
import Image from 'next/image'
import styles from '../styles/Home.module.css'
import { useState, useEffect } from 'react'
import 'tailwindcss/tailwind.css'

export default function Home() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  useEffect(() => {
    getData()
  }, []); 

  console.log(data?.data)
  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
      const data = await res.json();
      setData(data);
    } 
    catch (error) {
      setError(error);
    }
  }

  return (
    <div class="grid h-screen place-items-center">
      <div>
      {error && <div>Failed to load {error.toString()}</div>}
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
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send `, {
        method:'POST',
        body: JSON.stringify(body)
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
      <h1 class="font-mono text-xl">Selamat datang</h1>
      <p class="font-mono">Silahkan input data</p>
      <br />
      <form class="w-full max-w-sm border-2 rounded-md border-white-500" onSubmit={handleSubmit} id="form1">
        <div class="flex items-center border-b border-teal-500 py-2">
          <input class="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 px-2 leading-tight focus:outline-none" type="text" placeholder="Jane Doe" aria-label="text" name='data' />
          <button class="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded" type="submit">
          Submit
        </button>
        <button class="flex-shrink-0 border-transparent border-4 text-teal-500 hover:text-teal-800 text-sm py-1 px-2 rounded" type="button">
          Cancel
        </button>
      </div>
      </form>
      {error && <p>error: {error.toString()}</p>}
      {data && <p>success: {data}</p>}
      <hr />
      <br />
    </div>
  )
}

