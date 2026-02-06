import React, { useEffect } from 'react';
import { Form, Input, Button } from 'antd';

const SlotFiller = ({ slotInfo, pendingSlots, onSubmit }) => {
  const [form] = Form.useForm();

  // 防御性处理
  if (!slotInfo || !slotInfo.slots) return null;

  useEffect(() => {
    if (pendingSlots && slotInfo && slotInfo.slots) {
      const defaults = {};
      Object.entries(slotInfo.slots)
        .filter(([key]) => pendingSlots.includes(key))
        .forEach(([key, slot]) => {
          defaults[key] = slot.example;
        });
      form.resetFields();
      form.setFieldsValue(defaults);
    }
  }, [pendingSlots, slotInfo, form]);

  const handleFinish = values => {
    if (onSubmit) onSubmit(values);
  };

  return (
    <Form form={form} layout="inline" onFinish={handleFinish} style={{ marginBottom: 16 }}>
      {Object.entries(slotInfo.slots)
        .filter(([key]) => pendingSlots.includes(key))
        .map(([key, slot]) => (
          <Form.Item
            key={key}
            name={key}
            label={<span style={{ fontWeight: 500 }}>{slot.name}</span>}
            rules={[{ required: slot.required, message: '必填' }]}
          >
            <Input 
              placeholder={`请输入${slot.name}`} 
              style={{ borderRadius: 8 }}
            />
          </Form.Item>
        ))}
      <Form.Item>
        <Button 
          type="primary" 
          htmlType="submit"
          style={{ borderRadius: 8, background: 'linear-gradient(135deg, #3b82f6 0%, #2563eb 100%)' }}
        >
          补全信息
        </Button>
      </Form.Item>
    </Form>
  );
};

export default SlotFiller; 