import React, { FC, useState, useReducer } from "react";
import { Button, Form, Input, Modal } from "antd";
import axios from "axios";

import { endPoint } from "../../constants";
import { DataType } from "../../App"



type Props = {
    setUsers: ( users: any ) => void
}

export const AddedUrlForm:FC<Props> = ({ setUsers }) => {
    const [isVisibleModal, setIsVisibleModal] = useState(false)

    const showModal = () => {
        setIsVisibleModal(true);
      };
    
      const handleOk = () => {
        setIsVisibleModal(false);
      };
    
    const onFinish = (values: any) => {
        axios.post(`${endPoint}/added`, values).then(newList => {
            setUsers(newList.data)
        })
    }

    return (
        <>
        <Button type="primary" onClick={showModal}>
        Added user
      </Button>
        <Modal 
            title="Added user" 
            visible={isVisibleModal}
            onOk={handleOk}
            onCancel={handleOk}
            okText={"Close"}
            >
            <Form
            name="basic"
            labelCol={{ span: 8 }}
            wrapperCol={{ span: 16 }}
            initialValues={{ remember: true }}
            onFinish={onFinish}
            autoComplete="off"
        >
            <Form.Item
                label="Username"
                name="username"
                rules={[{ required: true, message: 'Please input your username!' }]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                label="Url"
                name="url"
                rules={[{ required: true, message: 'Please input your url!' }]}
            >
                <Input />
            </Form.Item>

            <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                <Button type="primary" htmlType="submit">
                    Added
                </Button>
            </Form.Item>
        </Form>
        </Modal>
        </>
    )

}