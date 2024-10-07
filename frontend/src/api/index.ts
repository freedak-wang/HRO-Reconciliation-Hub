import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080/api',
});

export const generateMonthlyBill = (contractId: number, billMonth: string) =>
  api.post('/bills', { contract_id: contractId, bill_month: billMonth });

export const calculateDifferences = (billId: number) =>
  api.get(`/bills/${billId}/differences`);

export const applyAdjustments = (billId: number, adjustments: any[]) =>
  api.post(`/bills/${billId}/adjustments`, { adjustments });

export default api;