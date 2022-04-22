import Head from 'next/head'
import Router from 'next/router'

export default function Home() {

  const routeChange = () =>{
    Router.push('/quiz') 
  }

  return (
    <div className='grid justify-center grid-rows-4 grid-flow-col gap-4'>
      <h1 className='font-black text-7xl'>Welcome to sneardsQuiz</h1>
      <h2 className='font-bold text-5xl ml-40 justify-center'>Are you ready to begin?</h2>
      <button onClick={routeChange} className='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'>Start quiz</button>
    </div>
  )
}
