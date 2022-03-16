import React, { useReducer, useEffect, useState } from 'react';
import { ListUrl } from "./components/ListUrl";
import { AddedUrlForm } from "./components/Form"
import { Layout } from "antd";
import axios from "axios";
import { endPoint } from "./constants"


export type DataType = {
    key: number
    username: string
    url: string
};

const App = () => {
    const [data, setData] = useState([])

    useEffect(() => {
        const req = axios.get(`${endPoint}/users`).then(res =>{
            setData(res.data)
        }) 
    }, [])

    const { Header, Content } = Layout;
    return (
        <Layout>
            <Header>
                <AddedUrlForm setUsers={setData}/>
            </Header>
            <Content>
                <ListUrl setUsers={setData} users={data}/>
            </Content>
        </Layout>
    )
}

export default App;
