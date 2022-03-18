import React, {FC, useEffect, useReducer, useState} from "react";
import {
    Table,
    Divider,
    Button,
    Row,
    Col
} from 'antd';
import "antd/dist/antd.css";
import axios from "axios";
import { endPoint } from "../../constants"
import { DataType } from "../../App"

const columns = [
    {
        title: 'Username',
        dataIndex: 'username',
        key: 'username',
    },
    {
        title: 'Url',
        dataIndex: 'url',
        key: 'url',
    }
];

type Props = {
    users: DataType[]
    setUsers: (users: any) => void

};

export const ListUrl:FC<Props>= ({ users, setUsers }) => {
    const [keySelected, setKeySelected] = useState<React.Key[] >([]);

    async function deleteUser () {
        if(keySelected.length === 1) {
            const [ key ] = keySelected 
            await axios.delete(`${endPoint}/user/${key}`, {
                headers: {
                    "Content-Type": "application/json",
                  }
            }).then((res) => {
                setUsers(res.data)
            })            
    }
};


    const rowSelection = {
        onChange: (selectedRowKeys: React.Key[]) => {
            setKeySelected(selectedRowKeys)
        },
        getCheckboxProps: (record: DataType) => ({
            disabled: record.username === 'Disabled User',
            username: record.username,
        }),
    };

    return (
        <Row gutter={16}>
            <Col span={12} offset={6}>
                {users && users.length &&
                    <Table
                        rowSelection={{
                            type: 'radio',
                            ...rowSelection
                        }}
                        columns={columns}
                        dataSource={users}
                        pagination={false}
                    />
                }
                <Divider/>
                <Button disabled={!keySelected.length} onClick={deleteUser} type="primary">Delete user</Button>
            </Col>

        </Row>
    );
};