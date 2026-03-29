#![allow(unused_comparisons)]

use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::hyperkd::hyperhv::interface::callbacks::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EventError {
    InvalidEventType,
    InvalidParameter,
    EventAlreadyExists,
    EventNotFound,
    NotInitialized,
    InsufficientResources,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EventState {
    Enabled,
    Disabled,
    Triggered,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EventType {
    Breakpoint,
    Exception,
    MemoryAccess,
    Syscall,
    Process,
    Thread,
    Module,
    Custom(u32),
}

#[derive(Debug, Clone, Copy)]
pub struct EventOptions {
    pub process_id: u32,
    pub thread_id: u32,
    pub core_id: u32,
    pub optional_param1: u64,
    pub optional_param2: u64,
    pub optional_param3: u64,
    pub optional_param4: u64,
}

impl Default for EventOptions {
    fn default() -> Self {
        Self {
            process_id: 0,
            thread_id: 0,
            core_id: 0,
            optional_param1: 0,
            optional_param2: 0,
            optional_param3: 0,
            optional_param4: 0,
        }
    }
}

#[derive(Debug, Clone)]
pub struct Event {
    pub id: u64,
    pub event_type: EventType,
    pub state: EventState,
    pub options: EventOptions,
    pub tag: u32,
    pub hit_count: u64,
    pub enabled: bool,
}

impl Event {
    pub fn new(id: u64, event_type: EventType, options: EventOptions, tag: u32) -> Self {
        Self {
            id,
            event_type,
            state: EventState::Enabled,
            options,
            tag,
            hit_count: 0,
            enabled: true,
        }
    }

    pub fn enable(&mut self) {
        self.state = EventState::Enabled;
        self.enabled = true;
    }

    pub fn disable(&mut self) {
        self.state = EventState::Disabled;
        self.enabled = false;
    }

    pub fn trigger(&mut self) {
        self.state = EventState::Triggered;
        self.hit_count += 1;
    }

    pub fn reset(&mut self) {
        self.state = EventState::Enabled;
        self.hit_count = 0;
    }
}

#[derive(Debug, Clone)]
pub struct EventResult {
    pub is_successful: bool,
    pub error: EventError,
    pub additional_info: u64,
}

impl EventResult {
    pub fn success() -> Self {
        Self {
            is_successful: true,
            error: EventError::InvalidEventType,
            additional_info: 0,
        }
    }

    pub fn failure(error: EventError) -> Self {
        Self {
            is_successful: false,
            error,
            additional_info: 0,
        }
    }

    pub fn failure_with_info(error: EventError, info: u64) -> Self {
        Self {
            is_successful: false,
            error,
            additional_info: info,
        }
    }
}

pub struct EventManager {
    events: alloc::collections::BTreeMap<u64, Event>,
    next_event_id: AtomicU64,
    initialized: AtomicBool,
    max_events: usize,
    event_callbacks: alloc::vec::Vec<unsafe fn(*mut Event)>,
}

impl EventManager {
    pub const fn new(max_events: usize) -> Self {
        Self {
            events: alloc::collections::BTreeMap::new(),
            next_event_id: AtomicU64::new(1),
            initialized: AtomicBool::new(false),
            max_events,
            event_callbacks: alloc::vec::Vec::new(),
        }
    }

    pub fn initialize(&mut self) -> Result<(), EventError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(EventError::NotInitialized);
        }

        self.events.clear();
        self.next_event_id.store(1, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.events.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn create_event(
        &mut self,
        event_type: EventType,
        options: EventOptions,
        tag: u32,
    ) -> Result<u64, EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        if self.events.len() >= self.max_events {
            return Err(EventError::InsufficientResources);
        }

        let event_id = self.next_event_id.fetch_add(1, Ordering::AcqRel);
        let event = Event::new(event_id, event_type, options, tag);

        self.events.insert(event_id, event);

        Ok(event_id)
    }

    pub fn remove_event(&mut self, event_id: u64) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        if !self.events.contains_key(&event_id) {
            return Err(EventError::EventNotFound);
        }

        self.events.remove(&event_id);
        Ok(())
    }

    pub fn enable_event(&mut self, event_id: u64) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        let event = self.events.get_mut(&event_id)
            .ok_or(EventError::EventNotFound)?;

        event.enable();
        Ok(())
    }

    pub fn disable_event(&mut self, event_id: u64) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        let event = self.events.get_mut(&event_id)
            .ok_or(EventError::EventNotFound)?;

        event.disable();
        Ok(())
    }

    pub fn trigger_event(&mut self, event_id: u64) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        let event = self.events.get_mut(&event_id)
            .ok_or(EventError::EventNotFound)?;

        event.trigger();

        for callback in &self.event_callbacks {
            unsafe {
                callback(event as *const Event as *mut Event);
            }
        }

        Ok(())
    }

    pub fn reset_event(&mut self, event_id: u64) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        let event = self.events.get_mut(&event_id)
            .ok_or(EventError::EventNotFound)?;

        event.reset();
        Ok(())
    }

    pub fn get_event(&self, event_id: u64) -> Option<&Event> {
        self.events.get(&event_id)
    }

    pub fn get_events(&self) -> alloc::vec::Vec<&Event> {
        self.events.values().collect()
    }

    pub fn get_enabled_events(&self) -> alloc::vec::Vec<&Event> {
        self.events.values()
            .filter(|e| e.enabled)
            .collect()
    }

    pub fn get_events_by_type(&self, event_type: EventType) -> alloc::vec::Vec<&Event> {
        self.events.values()
            .filter(|e| e.event_type == event_type)
            .collect()
    }

    pub fn get_events_by_tag(&self, tag: u32) -> alloc::vec::Vec<&Event> {
        self.events.values()
            .filter(|e| e.tag == tag)
            .collect()
    }

    pub fn get_event_count(&self) -> usize {
        self.events.len()
    }

    pub fn get_enabled_event_count(&self) -> usize {
        self.events.values()
            .filter(|e| e.enabled)
            .count()
    }

    pub fn register_event_callback(&mut self, callback: unsafe fn(*mut Event)) {
        self.event_callbacks.push(callback);
    }

    pub fn unregister_event_callback(&mut self, callback: unsafe fn(*mut Event)) {
        #[allow(unpredictable_function_pointer_comparisons)]
        {
            self.event_callbacks.retain(|&cb| cb != callback);
        }
    }

    pub fn apply_event(&mut self, event: &Event) -> EventResult {
        if !self.is_initialized() {
            return EventResult::failure(EventError::NotInitialized);
        }

        match event.event_type {
            EventType::Breakpoint => {
                unsafe {
                    if let Some(callback) = CALLBACKS.lock().trigger_events {
                        let mut post_event_required = false;
                        let mut regs = GuestRegs::default();
                        
                        let status = callback(
                            VmmEventType::Exception,
                            CallbackCallingStage::PreEvent,
                            core::ptr::null_mut(),
                            &mut post_event_required as *mut bool,
                            &mut regs as *mut GuestRegs,
                        );

                        if status == VmmCallbackStatus::Successful {
                            EventResult::success()
                        } else {
                            EventResult::failure(EventError::InvalidEventType)
                        }
                    } else {
                        EventResult::failure(EventError::InvalidEventType)
                    }
                }
            }
            EventType::Exception => {
                EventResult::success()
            }
            EventType::MemoryAccess => {
                EventResult::success()
            }
            EventType::Syscall => {
                EventResult::success()
            }
            EventType::Process => {
                EventResult::success()
            }
            EventType::Thread => {
                EventResult::success()
            }
            EventType::Module => {
                EventResult::success()
            }
            EventType::Custom(_) => {
                EventResult::success()
            }
        }
    }

    pub fn validate_event(&self, event: &Event) -> EventResult {
        if !self.is_initialized() {
            return EventResult::failure(EventError::NotInitialized);
        }

        match event.event_type {
            EventType::Breakpoint => {
                if event.options.optional_param1 == 0 {
                    return EventResult::failure_with_info(
                        EventError::InvalidParameter,
                        1,
                    );
                }
                EventResult::success()
            }
            EventType::Exception => {
                EventResult::success()
            }
            EventType::MemoryAccess => {
                if event.options.optional_param1 == 0 || 
                   event.options.optional_param2 == 0 {
                    return EventResult::failure_with_info(
                        EventError::InvalidParameter,
                        2,
                    );
                }
                EventResult::success()
            }
            EventType::Syscall => {
                EventResult::success()
            }
            EventType::Process => {
                EventResult::success()
            }
            EventType::Thread => {
                EventResult::success()
            }
            EventType::Module => {
                EventResult::success()
            }
            EventType::Custom(_) => {
                EventResult::success()
            }
        }
    }

    pub fn enable_all_events(&mut self) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        for event in self.events.values_mut() {
            event.enable();
        }

        Ok(())
    }

    pub fn disable_all_events(&mut self) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        for event in self.events.values_mut() {
            event.disable();
        }

        Ok(())
    }

    pub fn remove_all_events(&mut self) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        self.events.clear();
        Ok(())
    }

    pub fn reset_all_events(&mut self) -> Result<(), EventError> {
        if !self.is_initialized() {
            return Err(EventError::NotInitialized);
        }

        for event in self.events.values_mut() {
            event.reset();
        }

        Ok(())
    }
}

impl Drop for EventManager {
    fn drop(&mut self) {
        self.deinitialize();
    }
}

pub static EVENT_MANAGER: Mutex<EventManager> = Mutex::new(EventManager::new(1000));

pub fn initialize_event_manager() -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_event_manager() {
    let mut manager = EVENT_MANAGER.lock();
    manager.deinitialize();
}

pub fn create_event(
    event_type: EventType,
    options: EventOptions,
    tag: u32,
) -> Result<u64, EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.create_event(event_type, options, tag)
}

pub fn remove_event(event_id: u64) -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.remove_event(event_id)
}

pub fn enable_event(event_id: u64) -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.enable_event(event_id)
}

pub fn disable_event(event_id: u64) -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.disable_event(event_id)
}

pub fn trigger_event(event_id: u64) -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.trigger_event(event_id)
}

pub fn reset_event(event_id: u64) -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.reset_event(event_id)
}

pub fn get_event(event_id: u64) -> Option<Event> {
    let manager = EVENT_MANAGER.lock();
    manager.get_event(event_id).cloned()
}

pub fn get_events() -> alloc::vec::Vec<Event> {
    let manager = EVENT_MANAGER.lock();
    manager.get_events().into_iter().cloned().collect()
}

pub fn get_enabled_events() -> alloc::vec::Vec<Event> {
    let manager = EVENT_MANAGER.lock();
    manager.get_enabled_events().into_iter().cloned().collect()
}

pub fn get_events_by_type(event_type: EventType) -> alloc::vec::Vec<Event> {
    let manager = EVENT_MANAGER.lock();
    manager.get_events_by_type(event_type).into_iter().cloned().collect()
}

pub fn get_events_by_tag(tag: u32) -> alloc::vec::Vec<Event> {
    let manager = EVENT_MANAGER.lock();
    manager.get_events_by_tag(tag).into_iter().cloned().collect()
}

pub fn get_event_count() -> usize {
    let manager = EVENT_MANAGER.lock();
    manager.get_event_count()
}

pub fn get_enabled_event_count() -> usize {
    let manager = EVENT_MANAGER.lock();
    manager.get_enabled_event_count()
}

pub fn enable_all_events() -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.enable_all_events()
}

pub fn disable_all_events() -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.disable_all_events()
}

pub fn remove_all_events() -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.remove_all_events()
}

pub fn reset_all_events() -> Result<(), EventError> {
    let mut manager = EVENT_MANAGER.lock();
    manager.reset_all_events()
}

pub fn register_event_callback(callback: unsafe fn(*mut Event)) {
    let mut manager = EVENT_MANAGER.lock();
    manager.register_event_callback(callback)
}

pub fn unregister_event_callback(callback: unsafe fn(*mut Event)) {
    let mut manager = EVENT_MANAGER.lock();
    manager.unregister_event_callback(callback)
}

pub fn apply_event(event: &Event) -> EventResult {
    let mut manager = EVENT_MANAGER.lock();
    manager.apply_event(event)
}

pub fn validate_event(event: &Event) -> EventResult {
    let manager = EVENT_MANAGER.lock();
    manager.validate_event(event)
}
