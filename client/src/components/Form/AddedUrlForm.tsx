import { FC } from "react";
import { Button, Form, Input, Modal } from "antd";
import axios from "axios";

import { endPoint } from "../../constants";



type Props = {
    setUsers: ( users: any ) => void
}

export const AddedUrlForm:FC<Props> = ({ setUsers }) => {

    const onFinish = (values: any) => {
        axios.post(`${endPoint}/added`, values).then(newList => {
            setUsers(newList.data)
        })
    }

    function ModalForm() {
        Modal.info({
           title: "Добваить пользователя",
            content: (
                <>
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
                            label="Password"
                            name="password"
                            rules={[{ required: true, message: 'Please input your password!' }]}
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
                                Добавить
                            </Button>
                        </Form.Item>
                    </Form>
                </>
            ),
            closable: true,
            okText: "Завершить",
            width: "600px"
        })
    }

    return (
        <Button type="primary" onClick={ModalForm}>
            Добавить пользователя
        </Button>
    )

}