use booking_core::{
    errors::{BookingError, BookingResult},
    model::{Order, OrderId},
};
use crate::db::Database;

impl Database {
    pub async fn save_order(&self, order: &Order) -> BookingResult<()> {
        let mut tx = self.pool.begin().await.map_err(|e| {
            BookingError::RepositoryFailed {
                entity: "Transaction",
                reason: e.to_string(),
            }
        })?;

        let order_query = r#"
            INSERT INTO orders (id, user_id, total_amount, status, created_at)
            VALUES ($1, $2, $3, $4, $5)
        "#;

        sqlx::query(order_query)
            .bind(order.id.as_uuid())
            .bind(order.user_id.as_uuid())
            .bind(order.total_amount)
            .bind(serde_json::to_string(&order.status).unwrap())
            .bind(order.created_at)
            .execute(&mut *tx)
            .await
            .map_err(|e| BookingError::RepositoryFailed {
                entity: "Order",
                reason: e.to_string(),
            })?;

        let item_query = r#"
            INSERT INTO order_items (id, order_id, event_id, ticket_type_id, quantity, unit_price)
            VALUES ($1, $2, $3, $4, $5, $6)
        "#;

        for item in &order.items {
            sqlx::query(item_query)
                .bind(item.id)
                .bind(item.order_id.as_uuid())
                .bind(item.event_id.as_uuid())
                .bind(item.ticket_type_id.as_uuid())
                .bind(item.quantity)
                .bind(item.unit_price)
                .execute(&mut *tx)
                .await
                .map_err(|e| BookingError::RepositoryFailed {
                    entity: "OrderItem",
                    reason: e.to_string(),
                })?;
        }

        tx.commit().await.map_err(|e| BookingError::RepositoryFailed {
            entity: "Transaction Commit",
            reason: e.to_string(),
        })?;

        Ok(())
    }

    pub async fn order_exists(&self, id: &OrderId) -> BookingResult<bool> {
        let query = "SELECT EXISTS(SELECT 1 FROM orders WHERE id = $1)";
        
        let exists: (bool,) = sqlx::query_as(query)
            .bind(id.as_uuid())
            .fetch_one(&self.pool)
            .await
            .map_err(|e| BookingError::RepositoryFailed {
                entity: "Order Lookup",
                reason: e.to_string(),
            })?;

        Ok(exists.0)
    }
}