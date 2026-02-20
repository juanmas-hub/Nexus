use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};
use uuid::Uuid;
use std::fmt;

use crate::BookingError; 

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub struct OrderId(Uuid);

impl OrderId {
    pub fn new() -> Self {
        Self(Uuid::new_v4())
    }

    pub fn from_uuid(id: Uuid) -> Self {
        Self(id)
    }

    pub fn as_uuid(&self) -> Uuid {
        self.0
    }
}

impl fmt::Display for OrderId {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}", self.0)
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub struct UserId(Uuid);
impl UserId { pub fn from_uuid(id: Uuid) -> Self { Self(id) } }

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub struct EventId(Uuid);
impl EventId { pub fn from_uuid(id: Uuid) -> Self { Self(id) } }

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash, Serialize, Deserialize)]
pub struct TicketTypeId(Uuid);
impl TicketTypeId { pub fn from_uuid(id: Uuid) -> Self { Self(id) } }

#[derive(Debug, Clone, Copy, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "SCREAMING_SNAKE_CASE")]
pub enum OrderStatus {
    Pending,
    Confirmed,
    Cancelled,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct OrderItem {
    pub id: Uuid,
    pub order_id: OrderId,
    pub event_id: EventId,
    pub ticket_type_id: TicketTypeId,
    pub quantity: i32,
    pub unit_price: f64,
}

impl OrderItem {
    pub fn new(
        order_id: OrderId,
        event_id: EventId,
        ticket_type_id: TicketTypeId,
        quantity: i32,
        unit_price: f64,
    ) -> Result<Self, BookingError> {
        if quantity <= 0 {
            return Err(BookingError::ValidationFailed {
                field: "quantity",
                reason: "must be greater than 0",
            });
        }
        if unit_price < 0.0 {
            return Err(BookingError::ValidationFailed {
                field: "unit_price",
                reason: "cannot be negative",
            });
        }

        Ok(Self {
            id: Uuid::new_v4(),
            order_id,
            event_id,
            ticket_type_id,
            quantity,
            unit_price,
        })
    }
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Order {
    pub id: OrderId,
    pub user_id: UserId,
    pub total_amount: f64,
    pub status: OrderStatus,
    pub items: Vec<OrderItem>,
    pub created_at: DateTime<Utc>,
}

impl Order {
    pub fn new(user_id: UserId, items: Vec<OrderItem>) -> Result<Self, BookingError> {
        if items.is_empty() {
            return Err(BookingError::ValidationFailed {
                field: "items",
                reason: "order must contain at least one item",
            });
        }

        let total_amount = items
            .iter()
            .map(|item| item.unit_price * item.quantity as f64)
            .sum();

        Ok(Self {
            id: OrderId::new(),
            user_id,
            total_amount,
            status: OrderStatus::Pending,
            items,
            created_at: Utc::now(),
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn order_item_validates_quantity() {
        let result = OrderItem::new(
            OrderId::new(),
            EventId::from_uuid(Uuid::new_v4()),
            TicketTypeId::from_uuid(Uuid::new_v4()),
            0, // Invalid quantity
            100.0,
        );
        assert!(result.is_err());
    }

    #[test]
    fn order_validates_empty_items() {
        let result = Order::new(UserId::from_uuid(Uuid::new_v4()), vec![]);
        assert!(result.is_err());
    }
    
    #[test]
    fn order_calculates_total_correctly() {
        let order_id = OrderId::new();
        let item1 = OrderItem::new(order_id, EventId::from_uuid(Uuid::new_v4()), TicketTypeId::from_uuid(Uuid::new_v4()), 2, 50.0).unwrap();
        let item2 = OrderItem::new(order_id, EventId::from_uuid(Uuid::new_v4()), TicketTypeId::from_uuid(Uuid::new_v4()), 1, 150.0).unwrap();
        
        let order = Order::new(UserId::from_uuid(Uuid::new_v4()), vec![item1, item2]).unwrap();
        
        assert_eq!(order.total_amount, 250.0);
    }
}