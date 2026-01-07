"""Functions to manage a users shopping cart items."""


def add_item(current_cart, items_to_add):
    """Add items to shopping cart.

    :param current_cart: dict - the current shopping cart.
    :param items_to_add: iterable - items to add to the cart.
    :return: dict - the updated user cart dictionary.
    """

    for item in items_to_add:
        if item in current_cart:
            current_cart[item] += 1 
        else:
            current_cart[item] = 1
    return current_cart

def read_notes(notes):
    """Create user cart from an iterable notes entry.

    :param notes: iterable of items to add to cart.
    :return: dict - a user shopping cart dictionary.
    """
    item_list = {}
    for item in notes:
        item_list[item] =  item_list.get(item,0) + 1
    return item_list


def update_recipes(ideas, recipe_updates):
    """Update the recipe ideas dictionary.

    :param ideas: dict - The "recipe ideas" dict.
    :param recipe_updates: iterable -  with updates for the ideas section.
    :return: dict - updated "recipe ideas" dict.
    """
    ideas |= recipe_updates
    return ideas


def sort_entries(cart):
    """Sort a users shopping cart in alphabetically order.

    :param cart: dict - a users shopping cart dictionary.
    :return: dict - users shopping cart sorted in alphabetical order.
    """

    return dict(sorted(cart.items()))


def send_to_store(cart, aisle_mapping):
    """Combine users order to aisle and refrigeration information.

    :param cart: dict - users shopping cart dictionary.
    :param aisle_mapping: dict - aisle and refrigeration information dictionary.
    :return: dict - fulfillment dictionary ready to send to store.
    """
    fulfillment = {}
    for item in sorted(cart.keys(), reverse=True):
        quantity = cart[item]
        aisle, refrigerated = aisle_mapping[item]
        fulfillment[item] = [quantity, aisle, refrigerated]

    return fulfillment



def update_store_inventory(fulfillment_cart, store_inventory):
    """Update store inventory based on a fulfillment cart.

    :param fulfillment_cart: dict - items ordered with quantity, aisle, refrigeration
    :param store_inventory: dict - store inventory with quantity, aisle, refrigeration
    :return: dict - updated store inventory
    """
    for item, order_info in fulfillment_cart.items():
        ordered_qty = order_info[0]

        current_qty, aisle, refrigerated = store_inventory[item]
        new_qty = current_qty - ordered_qty

        if new_qty == 0:
            store_inventory[item] = ['Out of Stock', aisle, refrigerated]
        else:
            store_inventory[item] = [new_qty, aisle, refrigerated]

    return store_inventory

