import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import * as api from '../api';

export const generateMonthlyBill = createAsyncThunk(
  'bill/generate',
  async ({ contractId, billMonth }: { contractId: number; billMonth: string }) => {
    const response = await api.generateMonthlyBill(contractId, billMonth);
    return response.data;
  }
);

export const calculateDifferences = createAsyncThunk(
  'bill/calculateDifferences',
  async (billId: number) => {
    const response = await api.calculateDifferences(billId);
    return response.data;
  }
);

export const applyAdjustments = createAsyncThunk(
  'bill/applyAdjustments',
  async ({ billId, adjustments }: { billId: number; adjustments: any[] }) => {
    const response = await api.applyAdjustments(billId, adjustments);
    return response.data;
  }
);

const billSlice = createSlice({
  name: 'bill',
  initialState: {
    currentBill: null,
    differences: [],
    loading: false,
    error: null,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(generateMonthlyBill.pending, (state) => {
        state.loading = true;
      })
      .addCase(generateMonthlyBill.fulfilled, (state, action) => {
        state.loading = false;
        state.currentBill = action.payload;
      })
      .addCase(generateMonthlyBill.rejected, (state, action) => {
        state.loading = false;
        state.error = action.error.message;
      })
      // Add similar cases for calculateDifferences and applyAdjustments
  },
});

export default billSlice.reducer;